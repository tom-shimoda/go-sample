# >>> Register
# bindチェック
# ユーザー名だけ
# curl -X POST http://localhost:8000/api/register --data '{"username":"tarou"}'

# パスワードだけ
# curl -X POST http://localhost:8000/api/register --data '{"password":"passwd"}'

# ユーザー名、パスワード両方設定
# curl -X POST http://localhost:8000/api/register --data '{"username":"tarou", "password":"passwd"}'
# curl -X POST http://localhost:8000/api/register --data '{"username":"jirou", "password":"passwd"}'
# curl -X POST http://localhost:8000/api/register --data '{"username":"Hanako", "password":"passwd"}'
# curl -X POST http://localhost:8000/api/register --data '{"username":"hoge", "password":"12345678"}'


# >>> Login
# curl -X POST http://localhost:8000/api/login --data '{"username":"tarou", "password":"passwd"}'
# curl -X POST http://localhost:8000/api/login --data '{"username":"jirou", "password":"passwd"}'


# >>> jwt tokenを使用したAPIアクセス
# jqの引数について (https://www.tohoho-web.com/ex/jq.html#opt-compact-output)
#     -c: 改行やインデント無しのコンパクト形式で出力
#     -r: 出力データの文字列のダブルクォーテーションを取り除く
TOKEN=$(curl --silent -X POST http://localhost:8000/api/login --data '{"username":"tarou", "password":"passwd"}' | jq -rc .token)
# TOKEN=$(curl --silent -X POST http://localhost:8000/api/login --data '{"username":"jirou", "password":"passwd"}' | jq -rc .token)
# TOKEN=$(curl --silent -X POST http://localhost:8000/api/login --data '{"username":"HANAKO", "password":"passwd"}' | jq -rc .token)
# echo $TOKEN

# tokenを使わないアクセス
# curl --silent -X GET http://localhost:8000/api/admin/user
# tokenを使用したアクセス
curl --silent -X GET http://localhost:8000/api/admin/user -H "Authorization: Bearer ${TOKEN}"
