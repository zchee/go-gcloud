# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL = static
APP = go-gcloud
CMD_PREFIX = $(PKG)/cmd/
CMD = $(CMD_PREFIX)$(APP)

# ----------------------------------------------------------------------------
# target

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk

# ----------------------------------------------------------------------------
# overlays
