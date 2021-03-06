package route

import (
	"itflow/bug/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Level *xmux.GroupRoute

func init() {
	Level = xmux.NewGroupRoute("level")
	Level.Pattern("/level/get").Post(handle.LevelGet)
	Level.Pattern("/level/add").Post(handle.LevelAdd).End(midware.EndLog)
	Level.Pattern("/level/del").Get(handle.LevelDel).End(midware.EndLog)
	Level.Pattern("/level/update").Post(handle.LevelUpdate).End(midware.EndLog)
	Level.Pattern("/get/levels").Post(handle.GetLevels)
}
