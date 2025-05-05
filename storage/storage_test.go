package storage

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var testStorageService = &StorageService{}
var mockRedis *redis.Client
var mock redismock.ClientMock

func TestMain(m *testing.M) {
	mockRedis, mock = redismock.NewClientMock()

	testStorageService = &StorageService{
		redisClient: mockRedis,
	}

	m.Run()
}

func TestInitStorage(t *testing.T) {
	assert.NotNil(t, testStorageService.redisClient)
}

func TestSaveUrlMapping(t *testing.T) {
	mock.ExpectSet("shortUrl", "originalUrl", 24*time.Hour).SetVal("OK")

	testStorageService.SaveUrlMapping("shortUrl", "originalUrl", "userId")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOriginalUrl(t *testing.T) {
	mock.ExpectGet("shortUrl").SetVal("originalUrl")

	result, _ := testStorageService.GetOriginalUrl("shortUrl")

	assert.Equal(t, "originalUrl", result, "The returned URL should match the expected value")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOriginalUrl_KeyNotFound(t *testing.T) {
	mock.ExpectGet("shortUrl").RedisNil()

	result, err := testStorageService.GetOriginalUrl("shortUrl")

	assert.Equal(t, "", result, "The returned URL should be empty for a missing key")
	assert.Error(t, err, "An error should be returned when the key is not found")
	assert.EqualError(t, err, "Key not found: shortUrl", "The error message should match the expected value")

	assert.NoError(t, mock.ExpectationsWereMet())
}
