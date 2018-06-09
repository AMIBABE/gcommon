package errorcode

const (
	// UploadFail 图片上传失败, 将传入的图片格式改为正确的格式、适当的大小的图片放进消息体里面传输过来，如果传输仍然失败需要减小图片大小或者增加网络带宽进行尝试
	UploadFail = 3

	// AppCallLimited 应用调用次数超限, 调整程序合理调用API，等限频时间过了再调用，淘客的调用频率是系统按照上个月交易额自动修改的，修改后的频率会在官方论坛首页以公告形式通知，开发者可自行查看
	AppCallLimited = 7

	// HTTPActionNotAllowed HTTP方法被禁止, 请用大写的POST或GET，如果有图片等信息传入则一定要用POST才可以
	HTTPActionNotAllowed = 9

	// ServiceCurrentlyUnavailable 服务不可用, 多数是由未知异常引起的，仔细检查传入的参数是否符合文档描述
	ServiceCurrentlyUnavailable = 10

	// InsufficientISVPermissions 开发者权限不足
	InsufficientISVPermissions = 11

	// InsufficientUserPermissions 用户权限不足, 应用没有权限调用增值权限的接口
	InsufficientUserPermissions = 12

	// InsufficientPartnerPermissions 合作伙伴权限不足, 应用没有权限调用增值权限的接口
	InsufficientPartnerPermissions = 13

	// MissingMethod 缺少方法名参数, 传入的参数加入method字段
	MissingMethod = 21

	// InvalidMethod 不存在的方法名, 传入的method字段必需是你所调用的API的名称，并且该API是确实存在的
	InvalidMethod = 22

	// InvalidFormat 无效数据格式, 传入的format必需为json或xml中的一种
	InvalidFormat = 23

	// MissingSignature 缺少签名参数, 传入的参数中必需包含sign字段
	MissingSignature = 24

	// InvalidSignature 无效签名, 签名必需根据正确的算法算出来的.
	InvalidSignature = 25

	// MissingSession 缺少SessionKey参数, 传入的参数中必需包含session字段
	MissingSession = 26

	// InvalidSession 无效的SessionKey参数, 传入的session必需是用户绑定session拿到的，如果报session不合法可能是用户没有绑定session或session过期造成的，用户需要重新绑定一下然后传入新的sessionKey
	InvalidSession = 27

	// MissingAppKey 缺少AppKey参数, 传入的参数必需包含app_key字段
	MissingAppKey = 28

	// InvalidAppKey 无效的AppKey参数, 应用所处的环境跟选择的环境不一致，例如：应用处于沙箱测试环境，却选择在正式环境进行测试。
	InvalidAppKey = 29

	// MissingTimestamp 缺少时间戳参数, 传入的参数中必需包含timestamp参数
	MissingTimestamp = 30

	// InvalidTimestamp 非法的时间戳参数, 时间戳，格式为yyyy-mm-dd hh:mm:ss，例如：2008-01-25 20:23:30。淘宝API服务端允许客户端请求时间误差为10分钟
	InvalidTimestamp = 31

	// MissingVersion 缺少版本参数, 传入的参数中必需包含v字段
	MissingVersion = 32

	// InvalidVersion 非法的版本参数, 用户传入的版本号格式错误，必需为数字格式
	InvalidVersion = 33

	// UnsupportedVersion 不支持的版本号 , 用户传入的版本号没有被提供
	UnsupportedVersion = 34

	// InsufficientSessionPermissions 短授权权限不足, :simsun;">调用的是高危API
	InsufficientSessionPermissions = 42

	// ParameterError 参数错误, 一般是用户传入参数非法引起的，请仔细检查入参格式、范围是否一一对应
	ParameterError = 43

	// InvalidAccessToken 无效的access token
	InvalidAccessToken = 44

	// InvalidEncoding 编码错误, 一般是用户做http请求的时候没有用UTF-8编码请求造成的
	InvalidEncoding = 47
)

// ErrorMessageMap 定义常用的错误信息
var ErrorMessageMap map[int]string

func init() {
	ErrorMessageMap = make(map[int]string)
	ErrorMessageMap[UploadFail] = "图片上传失败, 将传入的图片格式改为正确的格式、适当的大小的图片放进消息体里面传输过来，如果传输仍然失败需要减小图片大小或者增加网络带宽进行尝试"
	ErrorMessageMap[AppCallLimited] = "应用调用次数超限, 调整程序合理调用API，等限频时间过了再调用，淘客的调用频率是系统按照上个月交易额自动修改的，修改后的频率会在官方论坛首页以公告形式通知，开发者可自行查看"
	ErrorMessageMap[HTTPActionNotAllowed] = "HTTP方法被禁止, 请用大写的POST或GET，如果有图片等信息传入则一定要用POST才可以"
	ErrorMessageMap[ServiceCurrentlyUnavailable] = "服务不可用, 多数是由未知异常引起的，仔细检查传入的参数是否符合文档描述"
	ErrorMessageMap[InsufficientISVPermissions] = "开发者权限不足"
	ErrorMessageMap[InsufficientUserPermissions] = "用户权限不足, 应用没有权限调用增值权限的接口"
	ErrorMessageMap[InsufficientPartnerPermissions] = "合作伙伴权限不足, 应用没有权限调用增值权限的接口"
	ErrorMessageMap[MissingMethod] = "缺少方法名参数, 传入的参数加入method字段"
	ErrorMessageMap[InvalidMethod] = "不存在的方法名, 传入的method字段必需是你所调用的API的名称，并且该API是确实存在的"
	ErrorMessageMap[InvalidFormat] = "无效数据格式, 传入的format必需为json或xml中的一种"
	ErrorMessageMap[MissingSignature] = "缺少签名参数, 传入的参数中必需包含sign字段"
	ErrorMessageMap[InvalidSignature] = "无效签名, 签名必需根据正确的算法算出来的."
	ErrorMessageMap[MissingSession] = "缺少SessionKey参数, 传入的参数中必需包含session字段"
	ErrorMessageMap[InvalidSession] = "无效的SessionKey参数, 传入的session必需是用户绑定session拿到的，如果报session不合法可能是用户没有绑定session或session过期造成的，用户需要重新绑定一下然后传入新的sessionKey"
	ErrorMessageMap[MissingAppKey] = "缺少AppKey参数, 传入的参数必需包含app_key字段"
	ErrorMessageMap[InvalidAppKey] = "无效的AppKey参数, 应用所处的环境跟选择的环境不一致，例如：应用处于沙箱测试环境，却选择在正式环境进行测试。"
	ErrorMessageMap[MissingTimestamp] = "缺少时间戳参数, 传入的参数中必需包含timestamp参数"
	ErrorMessageMap[InvalidTimestamp] = "非法的时间戳参数, 时间戳，格式为yyyy-mm-dd hh:mm:ss，例如：2008-01-25 20:23:30。淘宝API服务端允许客户端请求时间误差为10分钟"
	ErrorMessageMap[MissingVersion] = "缺少版本参数, 传入的参数中必需包含v字段"
	ErrorMessageMap[InvalidVersion] = "非法的版本参数, 用户传入的版本号格式错误，必需为数字格式"
	ErrorMessageMap[UnsupportedVersion] = "不支持的版本号 , 用户传入的版本号没有被提供"
	ErrorMessageMap[InsufficientSessionPermissions] = "短授权权限不足,调用的是高危API"
	ErrorMessageMap[ParameterError] = "参数错误, 一般是用户传入参数非法引起的，请仔细检查入参格式、范围是否一一对应"
	ErrorMessageMap[InvalidAccessToken] = "无效的access token"
	ErrorMessageMap[InvalidEncoding] = "编码错误, 一般是用户做http请求的时候没有用UTF-8编码请求造成的"
}
