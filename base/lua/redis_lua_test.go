package lua

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"git.shiyou.kingsoft.com/go/log"
	"github.com/go-redis/redis/v8"
	"testing"
)

type Score struct {
	Member string
	Score  float64
}

func (s *Score) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Score) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

type Member struct {
	Find        bool
	ID          string
	Rank        int64
	Score       float64
	PreRank     int64
	Leaderboard string
}

func TestRedisLua(t *testing.T) {
	InitRedisClient()
	scores := map[string]float64{
		"test_client_2": 22,
		"test_client_1": 11,
		"test_client_4": 44,
		"test_client_5": 55,
		"test_client_3": 33,
	}

	redisMembers := make([]any, 0, len(scores))

	for id, score := range scores {
		redisMembers = append(redisMembers, &Score{
			Member: id,
			Score:  score,
		})
	}

	script := redis.NewScript(updateMemberScoreAscLua)

	sha, err := script.Load(context.Background(), RedisClient).Result()
	if err != nil {
		log.Error("have err", "err", err)
		return
	}

	result, err := RedisClient.EvalSha(context.Background(), sha, []string{"LEADERBOARD:test"}, redisMembers...).Result()
	if err != nil {
		log.Error("have err", "err", err)
		return
	}

	resultArr, ok := result.([]any)
	if !ok {
		log.Error("ok have false err", "err")
		return
	}

	memberRanks, err := getResultRanks(scores, resultArr)
	if err != nil {
		return
	}

	for _, v := range memberRanks {
		fmt.Printf("memberRank %+v\n", v)
	}

	fmt.Println("================")

}

func getResultRanks(scores map[string]float64, resultArr []any) ([]*Member, error) {
	memberRanks := make([]*Member, 0, len(scores))

	for _, member := range resultArr {
		memberData, ok := member.([]any)
		if !ok {
			return nil, errors.New("member have an new err")
		}

		// if len(memberData) < minMembers {
		// 	return nil, fmt.Errorf(ErrMemberDataLengthWLength, len(memberData))
		// }

		id, ok := memberData[1].(string)
		if !ok {
			return nil, errors.New("member have an new err")
		}

		rank, ok := memberData[2].(int64)
		if !ok {
			return nil, errors.New("member have an new err")
		}

		preRank, ok := memberData[4].(int64)
		if !ok {
			return nil, errors.New("member have an new err")
		}

		memberRanks = append(memberRanks, &Member{
			ID:      id,
			Score:   scores[id],
			Rank:    rank,
			PreRank: preRank,
		})
	}
	return memberRanks, nil
}
