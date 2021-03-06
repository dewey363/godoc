package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/iiinsomnia/yiigo"
)

const gosessid = "GOSESSID"

var store = sessions.NewCookieStore([]byte(yiigo.EnvString("session", "secret", "N0awmAuS2OziVFu^9!*0LY7MeCRgQ&z0")))

// Get 获取session值
func Get(c *gin.Context, key string, defaultValule ...interface{}) (interface{}, error) {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		yiigo.Err(err.Error())
		return nil, err
	}

	var data interface{}

	// Get some session values.
	if v, ok := session.Values[key]; ok {
		data = v
	} else {
		if len(defaultValule) > 0 {
			data = defaultValule
		}
	}

	return data, nil
}

// Set 设置session值
func Set(c *gin.Context, key string, data interface{}, duration ...int) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	if len(duration) > 0 {
		session.Options = &sessions.Options{
			Path:   "/",
			MaxAge: duration[0],
		}
	}

	// Set some session values.
	session.Values[key] = data
	// Save it before we write to the response/return from the handler.
	err = session.Save(c.Request, c.Writer)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	return nil
}

// Delete 删除session值
func Delete(c *gin.Context, key string) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	delete(session.Values, key)

	err = session.Save(c.Request, c.Writer)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	return nil
}

// Destroy 删除session
func Destroy(c *gin.Context) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	err = session.Save(c.Request, c.Writer)

	if err != nil {
		yiigo.Err(err.Error())
		return err
	}

	return nil
}
