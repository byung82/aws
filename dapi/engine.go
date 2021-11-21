//go:generate go-enum -f=$GOFILE --marshal --lower --sql --sqlnullstr

package dapi

// DbEngine 크롤링 사이트
// ENUM(
// mysql
// postgres
//)
type DbEngine int
