package config

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"reflect"
	"strings"
	"sync"
	"unicode"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	debug      bool
	path       string
	configType string
	lock       sync.RWMutex

	// Method for creating data structure.
	newDataFunc func() any

	// Struct instance used to store config data.
	data any

	// all keys in data struct
	dataStructAllKeys []string

	viper   *viper.Viper
	watcher *viper.Viper

	md5 string

	// It is required that the data struct and the fields in the config file are strictly matched.
	exactDataStruct bool

	checkMissingKey bool

	// It is used to verify the correctness during config loading.
	verifyFunc func(any) error
}

func New(path string, newDataFunc func() any, options ...Option) (*Config, error) {
	slog.Debug("new config", "path", path)

	// check path
	if path == "" {
		return nil, errors.New("no file path")
	}
	// check newDataFunc
	if newDataFunc == nil {
		return nil, errors.New("no newDataFunc")
	}

	// config type
	str := strings.Split(path, ".")
	if len(str) <= 1 {
		return nil, errors.New("config type err")
	}
	configType := str[len(str)-1]

	ret := &Config{
		path:              path,
		configType:        configType,
		newDataFunc:       newDataFunc,
		data:              newDataFunc(),
		dataStructAllKeys: make([]string, 0, 32),
	}
	for _, op := range options {
		op(ret)
	}

	// Check the conflict between missing-key and proto
	if _, ok := ret.data.(proto.Message); ok && ret.checkMissingKey {
		return nil, errors.New("conflict between missing-key and proto")
	}

	// init allKeys
	t := reflect.TypeOf(ret.data)
	filedStruct("", t.Elem(), &ret.dataStructAllKeys)

	// read
	err := ret.read()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func watch(c *Config, reloadFunc func(any)) {
	if reloadFunc == nil {
		return
	}

	// new watch viper
	watcher := viper.New()
	replacer := strings.NewReplacer(".", "_")
	watcher.SetEnvKeyReplacer(replacer)
	watcher.AutomaticEnv()
	watcher.SetConfigFile(c.path)
	watcher.SetConfigType(c.configType)

	// watch
	watcher.WatchConfig()
	watcher.OnConfigChange(func(event fsnotify.Event) {
		slog.Info("OnConfigChange", "name", event.Name, "op", event.Op)
		if err := c.read(); err != nil {
			slog.Warn("config read err", "error", err)
			if err == ErrRepeatedRead {
				slog.Info("ErrRepeatedRead", "MD5", c.GetMD5())
			}
			return
		}
		reloadFunc(c.GetData())
	})

	c.watcher = watcher
}

var ErrMD5Value = errors.New("MD5 value err")

var ErrRepeatedRead = errors.New("repeated read")

func (c *Config) read() error {
	// compare newMD5
	newMD5, err := c.genMD5()
	if err != nil {
		return err
	}
	if newMD5 == "" {
		return ErrMD5Value
	}
	oldMD5 := c.GetMD5()
	if oldMD5 != "" && oldMD5 == newMD5 {
		return ErrRepeatedRead
	}

	// new viper
	replacer := strings.NewReplacer(".", "_")
	newViper := viper.New()
	newViper.SetEnvKeyReplacer(replacer)
	newViper.AutomaticEnv()

	// read config file
	newViper.SetConfigFile(c.path)
	newViper.SetConfigType(c.configType)
	if err := newViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		}
		return err
	}

	return c.update(newViper, newMD5)
}

func (c *Config) update(newViper *viper.Viper, newMD5 string) error {
	// unmarshal data
	var newData = c.newDataFunc()
	err := unmarshal(newViper, newData, c.exactDataStruct)
	if err != nil {
		return err
	}

	// check missing keys
	if c.checkMissingKey {
		err = c.checkMissingKeys(newViper)
		if err != nil {
			return err
		}
	}

	// verify
	if c.verifyFunc != nil {
		err := c.verifyFunc(newData)
		if err != nil {
			return err
		}
	}

	// update
	c.lock.Lock()
	defer c.lock.Unlock()
	c.viper = newViper
	c.data = newData
	c.md5 = newMD5
	return nil
}

func (c *Config) genMD5() (string, error) {
	bytes, err := os.ReadFile(c.path)
	if err != nil {
		return "", err
	}
	if c.debug {
		slog.Debug(string(bytes))
	}

	m := md5.New()
	if _, err := m.Write(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

func prefix(src string, allKeys []string) bool {
	lowSrc := strings.ToLower(src)
	for _, v := range allKeys {
		lowPrefix := strings.ToLower(v)
		if strings.HasPrefix(lowSrc, lowPrefix) {
			return true
		}
	}
	return false
}

func (c *Config) checkMissingKeys(viper *viper.Viper) error {
	missingKeys := ""

	// check config context missing-keys
	for _, k := range viper.AllKeys() {
		if viper.IsSet(k) {
			continue
		}
		missingKeys += " " + k
	}

	// check dataStruct missing-keys
	for _, k := range c.dataStructAllKeys {
		if !viper.IsSet(k) && !prefix(k, viper.AllKeys()) {
			missingKeys += " " + k
		}
	}

	// check result
	if missingKeys == "" {
		return nil
	}
	return errors.New("missing keys:" + missingKeys)
}

func filedStruct(prefix string, t reflect.Type, arr *[]string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		if len(name) != 0 && !unicode.IsUpper(int32([]byte(name)[0])) {
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			if prefix == "" {
				filedStruct(name, field.Type, arr)
			} else {
				filedStruct(prefix+"."+name, field.Type, arr)
			}
			continue
		} else if field.Type.Kind() == reflect.Ptr {
			if prefix == "" {
				filedStruct(name, field.Type.Elem(), arr)
			} else {
				filedStruct(prefix+"."+name, field.Type.Elem(), arr)
			}
			continue
		}

		if prefix == "" {
			*arr = append(*arr, name)
		} else {
			*arr = append(*arr, prefix+"."+name)
		}
	}
}

func unmarshal(viper *viper.Viper, v any, exactDataStruct bool) error {
	// to pb
	if m, ok := v.(proto.Message); ok {
		data, err := marshalJSON(convertMap(viper.AllSettings()))
		if err != nil {
			return err
		}
		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, m)
	}

	// to struct
	if exactDataStruct {
		return viper.UnmarshalExact(v)
	}
	return viper.Unmarshal(v)
}

func convertMap(src any) any {
	switch m := src.(type) {
	case map[string]any:
		dst := make(map[string]any, len(m))
		for k, v := range m {
			dst[k] = convertMap(v)
		}
		return dst
	case map[any]any:
		dst := make(map[string]any, len(m))
		for k, v := range m {
			dst[fmt.Sprint(k)] = convertMap(v)
		}
		return dst
	case []any:
		dst := make([]any, len(m))
		for k, v := range m {
			dst[k] = convertMap(v)
		}
		return dst
	case []byte:
		// there will be no binary data in the config data
		return string(m)
	default:
		return src
	}
}

func marshalJSON(v any) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}

func (c *Config) GetViper() *viper.Viper {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.viper
}

func (c *Config) GetData() any {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.data
}

func (c *Config) GetMD5() string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.md5
}
