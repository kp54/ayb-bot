/*
 * Misskey API
 *
 * **Misskey is a decentralized microblogging platform.**   ## Usage **APIはすべてPOSTでリクエスト/レスポンスともにJSON形式です。** 一部のAPIはリクエストに認証情報(APIキー)が必要です。リクエストの際に`i`というパラメータでAPIキーを添付してください。  ### 自分のアカウントのAPIキーを取得する 「設定 > API」で、自分のAPIキーを取得できます。  > アカウントを不正利用される可能性があるため、このトークンは第三者に教えないでください(アプリなどにも入力しないでください)。  ### アプリケーションとしてAPIキーを取得する 直接ユーザーのAPIキーをアプリケーションが扱うのはセキュリティ上のリスクがあるので、 アプリケーションからAPIを利用する際には、アプリケーションとアプリケーションを利用するユーザーが結び付けられた専用のAPIキーを発行します。  #### 1.アプリケーションを登録する まず、あなたのアプリケーションやWebサービス(以後、あなたのアプリと呼びます)をMisskeyに登録します。 [デベロッパーセンター](/dev)にアクセスし、「アプリ > アプリ作成」からアプリを作成してください。  登録が済むとあなたのアプリのシークレットキーが入手できます。このシークレットキーは後で使用します。  > アプリに成りすまされる可能性があるため、極力このシークレットキーは公開しないようにしてください。</p>  #### 2.ユーザーに認証させる アプリを使ってもらうには、ユーザーにアカウントへのアクセスの許可をもらう必要があります。  認証セッションを開始するには、[https://misskey.kurume-nct.com/api/auth/session/generate](#operation/auth/session/generate) へパラメータに`appSecret`としてシークレットキーを含めたリクエストを送信します。 レスポンスとして認証セッションのトークンや認証フォームのURLが取得できるので、認証フォームのURLをブラウザで表示し、ユーザーにフォームを提示してください。  あなたのアプリがコールバックURLを設定している場合、 ユーザーがあなたのアプリの連携を許可すると設定しているコールバックURLに`token`という名前でセッションのトークンが含まれたクエリを付けてリダイレクトします。  あなたのアプリがコールバックURLを設定していない場合、ユーザーがあなたのアプリの連携を許可したことを(何らかの方法で(たとえばボタンを押させるなど))確認出来るようにしてください。  #### 3.アクセストークンを取得する ユーザーが連携を許可したら、[https://misskey.kurume-nct.com/api/auth/session/userkey](#operation/auth/session/userkey) へリクエストを送信します。  上手くいけば、認証したユーザーのアクセストークンがレスポンスとして取得できます。おめでとうございます！  アクセストークンが取得できたら、*「ユーザーのアクセストークン+あなたのアプリのシークレットキーをsha256したもの」*をAPIキーとして、APIにリクエストできます。  APIキーの生成方法を擬似コードで表すと次のようになります: ``` js const i = sha256(userToken + secretKey); ``` 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// DriveFolder struct for DriveFolder
type DriveFolder struct {
	// The unique identifier for this Drive folder.
	Id string `json:"id"`
	// The date that the Drive folder was created.
	CreatedAt time.Time `json:"createdAt"`
	// The folder name.
	Name string `json:"name"`
	// The count of child folders.
	FoldersCount float32 `json:"foldersCount,omitempty"`
	// The count of child files.
	FilesCount float32 `json:"filesCount,omitempty"`
	// The parent folder ID of this folder.
	ParentId *string `json:"parentId,omitempty"`
	Parent DriveFolder `json:"parent,omitempty"`
}
