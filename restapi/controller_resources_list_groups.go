package restapi

import (
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/go-openapi/runtime/middleware"
)

// ListResourceGroupsController provides a list of resource groups. This is an exploratory read-only endpoint.
func ListResourceGroupsController(ch *helpers.ContextHelper) resources.ListResourceGroupsHandlerFunc {
	return resources.ListResourceGroupsHandlerFunc(func(params resources.ListResourceGroupsParams) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		return resources.NewListResourceGroupsOK().WithPayload(&models.ResponseListResources{
			Links:    helpers.HalRootRscLinks(ch),
			Embedded: buildResourceGroupResponse([]string{"tf", "identity", "vcs"}, ch),
		})
	})
}
