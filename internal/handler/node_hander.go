package handler

import (
	"dropbox-clone/internal/form"
	"dropbox-clone/internal/service"
	util "dropbox-clone/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NodeHandler struct {
	NodeService service.NodeService
}

func (h *NodeHandler) ListingNodes(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDUint, _ := userID.(uint)

	ParentID, exist := c.Params.Get("id")
	var parentID *string
	var title string
	var header string
	if exist == true {
		parentID = &ParentID
		node, _ := h.NodeService.NodeRepository.GetNodeByID(*parentID)
		title = node.Name
		header = node.Name + " Files"
	} else {
		parentID = nil
		title = "Home"
		header = "All Files"
	}

	nodes, _ := h.NodeService.ListNode(userIDUint, parentID)
	var nodeResponses []*form.NodeResponse
	for _, node := range nodes {
		nodeResponses = append(nodeResponses, &form.NodeResponse{
			ID:        node.ID,
			ParentID:  node.ParentID,
			Type:      string(node.Type),
			Name:      node.Name,
			FileSize:  node.FileSize,
			FileType:  node.FileType,
			CreatedAt: util.FormatDateTime(node.CreatedAt),
			UpdatedAt: util.FormatDateTime(node.UpdatedAt),
		})
	}
	util.RenderLayout(c, "index.templ", gin.H{"Title": title, "Header": header, "ParentID": parentID, "Nodes": nodeResponses})
}
func (h *NodeHandler) CreateNode(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDUint, _ := userID.(uint)
	Name := c.PostForm("Name")
	Type := c.PostForm("Type")

	ParentID := c.PostForm("ParentID")
	var parentID *string
	if ParentID != "" {
		parentID = &ParentID
	} else {
		parentID = nil
	}

	node, err := h.NodeService.CreateNode(Type, userIDUint, Name, util.StringPtr(*parentID))
	if err != nil {
		util.RespondWithError(c, http.StatusBadRequest, err.Error(), parentID)
		return
	}

	util.RespondWithSuccess(c, "Node Created", gin.H{
		"ID":   node.ID,
		"Name": node.Name,
	})
}

func (h *NodeHandler) DeleteNode(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDUint, _ := userID.(uint)

	ID, err := c.Params.Get("id")
	if err == false {
		util.RespondWithError(c, http.StatusBadRequest, "ID not found", nil)
		return
	}

	status, nodeErr := h.NodeService.DeleteNode(userIDUint, ID)
	if nodeErr != nil {
		util.RespondWithError(c, http.StatusBadRequest, nodeErr.Error(), nil)
		return
	}

	util.RespondWithSuccess(c, "Node Deleted", gin.H{"status": status})
}
