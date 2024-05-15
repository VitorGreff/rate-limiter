package algorithms

import (
	"rate-limiter/models"

	"github.com/labstack/echo/v4"
)

func TokenBucket(c echo.Context, buckets *[]models.Bucket) error {
	var (
		index int
		cond  bool
	)

	requestIpAddr := c.RealIP()
	if cond, index = models.BucketExist(*buckets, requestIpAddr); cond == false {
		*buckets = append(*buckets, *models.BuildBucket(requestIpAddr, 1.0))
		index = len(*buckets) - 1
	}

	err := (*buckets)[index].TakeToken()
	if err != nil {
    return err
	}
	return nil
}
