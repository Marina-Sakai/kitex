include "sonic_base.thrift"
namespace go kitex.test.server

enum TestEnum {
    FIRST = 1,
    SECOND = 2,
    THIRD = 3,
    FOURTH = 4,
}

struct GenericRequest {
    1: i64 vint64(api.query = 'vint64', api.vd = "$>0&&$<200")
    2: string text(api.body = 'text')
    3: list<i32> nums(api.body = 'nums')
    4: string token(api.header = 'token')
    5: list<string> items(api.query = 'items')
    6: i32 version(api.path = 'version')
    255: optional sonic_base.Base Base
}

struct GenericResponse {
    1: i64 vint64(api.header = 'vint64', api.vd = "$>0&&$<200")
    2: string text(api.body = 'text')
    3: list<i32> nums(api.body = 'nums')
    4: string token(api.header = 'token')
    5: list<string> items(api.header = 'items')
    6: i32 version(api.header = 'version')
    7: i32 http_code(api.http_code = '')
    8: optional list<string> boo = ["far", "boo"](api.header = 'boo')
    9: optional TestEnum tenum = TestEnum.THIRD(api.body = 'test_enum')
    10: optional list<sonic_base.BaseElem> baseElems = [{"a": 34}, {"a": 56}](api.body = 'base_elems' go.tag='json:\"base_elems,omitempty\"')
    11: optional set<string> keys = ["aaa", "bbb"](api.body = 'keys')
    12: optional map<i32, sonic_base.BaseElem> values = {
            12: {
                "a": 34
            }
        }(api.body = 'values')
    255: optional sonic_base.BaseResp BaseResp
}

service STService {
    GenericResponse testGeneric(1: GenericRequest req)(api.post = '/life/client/:version', api.baseurl = 'example.com', api.param = 'true', api.serializer = 'json')
}