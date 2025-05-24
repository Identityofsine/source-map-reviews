package cookies

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cookies struct {
	context *gin.Context
}

func NewCookies(c *gin.Context) *Cookies {
	return &Cookies{
		context: c,
	}
}

func (c *Cookies) Get(name string) (string, error) {
	value, err := c.context.Cookie(name)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cookies) GetInt(name string) (int, error) {

	value, err := c.Get(name)
	if err != nil {
		return 0, err
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func (c *Cookies) Set(name, value string, maxAge int) error {
	c.context.SetCookie(name, value, maxAge, "/", "", false, true)
	return nil
}

func (c *Cookies) SetInt(name string, value int64, maxAge int) error {
	strValue := strconv.FormatInt(value, 10)
	return c.Set(name, strValue, maxAge)
}

func (c *Cookies) SetIfNotExists(name, value string, maxAge int) error {
	// Check if the cookie already exists
	if _, err := c.Get(name); err == nil {
		return nil // Cookie already exists, do nothing
	}
	// Set the cookie if it does not exist
	return c.Set(name, value, maxAge)
}

func (c *Cookies) Delete(name string) error {
	c.context.SetCookie(name, "", -1, "/", "", false, true)
	return nil
}
