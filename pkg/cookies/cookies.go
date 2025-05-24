package cookies

import "github.com/gin-gonic/gin"

type Cookies struct {
	context *gin.Context
}

func newCookies(c *gin.Context) *Cookies {
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

func (c *Cookies) Set(name, value string, maxAge int) error {
	c.context.SetCookie(name, value, maxAge, "/", "", false, true)
	return nil
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
