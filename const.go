package gosphinx

const (
	// known searchd commands
	SEARCHD_COMMAND_SEARCH     = 0
	SEARCHD_COMMAND_EXCERPT    = 1
	SEARCHD_COMMAND_UPDATE     = 2
	SEARCHD_COMMAND_KEYWORDS   = 3
	SEARCHD_COMMAND_PERSIST    = 4
	SEARCHD_COMMAND_STATUS     = 5
	SEARCHD_COMMAND_FLUSHATTRS = 7

	// current client-side command implementation versions
	VER_MAJOR_PROTO        = 0x1
	VER_COMMAND_SEARCH     = 0x120
	VER_COMMAND_EXCERPT    = 0x104
	VER_COMMAND_UPDATE     = 0x103
	VER_COMMAND_KEYWORDS   = 0x100
	VER_COMMAND_STATUS     = 0x101
	VER_COMMAND_FLUSHATTRS = 0x100

	// known searchd status codes
	SEARCHD_OK      = 0
	SEARCHD_ERROR   = 1
	SEARCHD_RETRY   = 2
	SEARCHD_WARNING = 3

	// known ranking modes (extended2 mode only)
	SPH_RANK_PROXIMITY_BM15 = 0 // default mode, phrase proximity major factor and BM15 minor one
	SPH_RANK_BM15           = 1 // statistical mode, BM15 ranking only (faster but worse quality)
	SPH_RANK_NONE           = 2 // no ranking, all matches get a weight of 1
	SPH_RANK_WORDCOUNT      = 3 // simple word-count weighting, rank is a weighted sum of per-field keyword occurence counts
	SPH_RANK_PROXIMITY      = 4
	SPH_RANK_MATCHANY       = 5
	SPH_RANK_FIELDMASK      = 6
	SPH_RANK_SPH04          = 7
	SPH_RANK_EXPR           = 8
	SPH_RANK_TOTAL          = 9

	// aliases; to be retired
	SPH_RANK_PROXIMITY_BM25 = 0
	SPH_RANK_BM25           = 1

	// known sort modes
	SPH_SORT_RELEVANCE     = 0
	SPH_SORT_ATTR_DESC     = 1
	SPH_SORT_ATTR_ASC      = 2
	SPH_SORT_TIME_SEGMENTS = 3
	SPH_SORT_EXTENDED      = 4

	// known filter types
	SPH_FILTER_VALUES      = 0
	SPH_FILTER_RANGE       = 1
	SPH_FILTER_FLOATRANGE  = 2
	SPH_FILTER_STRING      = 3
	SPH_FILTER_STRING_LIST = 6

	// known attribute types
	SPH_ATTR_NONE    = 0
	SPH_ATTR_INTEGER = 1
	SPH_ATTR_BOOL    = 4
	SPH_ATTR_FLOAT   = 5
	SPH_ATTR_BIGINT  = 6
	SPH_ATTR_STRING  = 7
	SPH_ATTR_FACTORS = 1001
	SPH_ATTR_MULTI   = 0x40000001
	SPH_ATTR_MULTI64 = 0x40000002

	// known grouping functions
	SPH_GROUPBY_DAY      = 0
	SPH_GROUPBY_WEEK     = 1
	SPH_GROUPBY_MONTH    = 2
	SPH_GROUPBY_YEAR     = 3
	SPH_GROUPBY_ATTR     = 4
	SPH_GROUPBY_ATTRPAIR = 5

	// known match modes
	SPH_MATCH_ALL       = 0
	SPH_MATCH_ANY       = 1
	SPH_MATCH_PHRASE    = 2
	SPH_MATCH_BOOLEAN   = 3
	SPH_MATCH_EXTENDED  = 4
	SPH_MATCH_FULLSCAN  = 5
	SPH_MATCH_EXTENDED2 = 6
)
