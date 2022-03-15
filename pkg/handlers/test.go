 package handlers
import(
  "net/http"
//  "wxcrm/pkg/backend"
  "github.com/gin-gonic/gin"
)

func (h *HandlerVar)UserTest(c *gin.Context) {
// c.Header("Access-Control-Allow-Origin","http://vs5gqq.natappfree.cc/")
  dmaiuser := h.DB.ViewUser("wangbo")
  c.SetCookie("wxcrm_username", dmaiuser.Userid, 7200, "/", h.Opts.ServerName, false, true)
  c.JSON(http.StatusOK,gin.H{"status": "200","CookieName": "wxcrm_username","Userinfo":dmaiuser,"Test": "true"})
}


