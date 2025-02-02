/*
* Large-Scale Discovery, a network scanning solution for information gathering in large IT/OT network environments.
*
* Copyright (c) Siemens AG, 2016-2024.
*
* This work is licensed under the terms of the MIT license. For a copy, see the LICENSE file in the top-level
* directory or visit <https://opensource.org/licenses/MIT>.
*
 */

package core

import (
	"github.com/gin-gonic/gin"
	scanUtils "github.com/siemens/GoScans/utils"
	"github.com/siemens/Large-Scale-Discovery/web_backend/database"
)

const contextKey = "storage"

// ContextStorage contains data relevant throughout the lifetime of a request
type ContextStorage struct {
	Logger      scanUtils.Logger
	CurrentUser *database.T_user
}

// SetContextStorage allows to set the context storage
func SetContextStorage(ctx *gin.Context, c *ContextStorage) {
	ctx.Set(contextKey, c)
}

// GetContextLogger retrieves a reference to the tagged logger of the current requests context. This allows to use
// the same tagged logger within one request, across multiple handlers and functions.
func GetContextLogger(ctx *gin.Context) scanUtils.Logger {
	return getContextStorage(ctx).Logger
}

// GetContextUser retrieves a reference to the currently authenticated user from the requests context, which is set
// after successful authentication. Referenced values can also be updated throughout a request context.
func GetContextUser(ctx *gin.Context) *database.T_user {
	return getContextStorage(ctx).CurrentUser
}

// UnsetContextUser removes the user from the context store for further processing. This function is only required in
// the logout handler to prevent the core.Respond() function from automatically re-generating and returning a new
// authentication token
func UnsetContextUser(ctx *gin.Context) {
	contextStorage := getContextStorage(ctx)
	contextStorage.CurrentUser = nil
}

// getContextStorage retrieves a reference to the complete context storage from the current request context.
func getContextStorage(ctx *gin.Context) *ContextStorage {
	return ctx.Value(contextKey).(*ContextStorage)
}
