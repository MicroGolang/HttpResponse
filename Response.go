/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Thursday 26 September 2019 - 11:06:52
** @Filename:				response.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Thursday 26 September 2019 - 11:08:20
*******************************************************************************/

package		response

import		"encoding/json"
import		"net/http"
import		"time"

/*Resolve *********************************************************************
*	Send a response to the client
******************************************************************************/
func	Resolve(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

/*ResolveWithError ************************************************************
*	Send an error response to the client
******************************************************************************/
func	ResolveWithError(w http.ResponseWriter, err error, status int) {
	if (err != nil) {
		http.Error(w, err.Error(), status)
		return
	}
	http.Error(w, `false`, status)
}

/*ResolveCookie ***************************************************************
*	Add a cookie for the next response. Default expiration is 2 months
******************************************************************************/
func	ResolveCookie(w *http.ResponseWriter, cookieName, cookieValue string) {
	Cookie := &http.Cookie {
		Name:		cookieName,
		Value:		cookieValue,
		Path:		`/`,
		Expires:    time.Now().Add(time.Hour * 24 * 30 * 2),
		RawExpires:	time.Now().Add(time.Hour * 24 * 30 * 2).String(),
		HttpOnly:	true,
	}
	http.SetCookie(*w, Cookie)
}