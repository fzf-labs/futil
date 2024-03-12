package httputil

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client := NewClient()
	response, err := client.R().SetContext(ctx).Post("http://www.baidu.com")
	fmt.Println(response.Status)
	fmt.Println(err)
	time.Sleep(time.Second * 10)
	assert.Equal(t, nil, err)
}
