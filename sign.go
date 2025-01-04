package mobile_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"
)

func GenerateTP1Sign(appid, appSecret, timestamp string) string {
	signatureString := fmt.Sprintf("appid=%s&app_secret=%s&times=%s", appid, appSecret, timestamp)
	hash := md5.New()
	hash.Write([]byte(signatureString))
	hashString := hex.EncodeToString(hash.Sum(nil))
	signature := strings.ToUpper(hashString)
	return signature
}

func GenerateTP2Sign(channelID, apiKey string) string {
	signatureString := fmt.Sprintf("channel_id=%s&apikey=%s", channelID, apiKey)
	hash := md5.New()
	hash.Write([]byte(signatureString))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
func GetTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// TransIDGenerator is responsible for generating unique transaction IDs.
type TransIDGenerator struct {
	appid    string
	sequence int64 // 8位数字序列，从1开始递增
	mu       sync.Mutex
}

// NewTransIDGenerator creates a new instance of TransIDGenerator.
func NewTransIDGenerator(appid string) *TransIDGenerator {
	return &TransIDGenerator{
		appid:    appid,
		sequence: 1, // 从1开始递增
	}
}

// Generate generates a unique transaction ID based on the given rules.
func (g *TransIDGenerator) Generate() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 获取当前时间戳，格式为 YYYYMMDDHHMISS
	timestamp := time.Now().Format("20060102150405")

	// 生成8位数字序列，不足8位的前面补0
	sequenceStr := fmt.Sprintf("%08d", g.sequence)
	g.sequence++ // 序列号递增

	// 拼接 APPID + 时间戳 + 序列号
	transid := fmt.Sprintf("%s%s%s", g.appid, timestamp, sequenceStr)

	return transid
}
