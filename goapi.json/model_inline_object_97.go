/*
 * Misskey API
 *
 * **Misskey is a decentralized microblogging platform.**   ## Usage **APIはすべてPOSTでリクエスト/レスポンスともにJSON形式です。** 一部のAPIはリクエストに認証情報(APIキー)が必要です。リクエストの際に`i`というパラメータでAPIキーを添付してください。  ### 自分のアカウントのAPIキーを取得する 「設定 > API」で、自分のAPIキーを取得できます。  > アカウントを不正利用される可能性があるため、このトークンは第三者に教えないでください(アプリなどにも入力しないでください)。  ### アプリケーションとしてAPIキーを取得する 直接ユーザーのAPIキーをアプリケーションが扱うのはセキュリティ上のリスクがあるので、 アプリケーションからAPIを利用する際には、アプリケーションとアプリケーションを利用するユーザーが結び付けられた専用のAPIキーを発行します。  #### 1.アプリケーションを登録する まず、あなたのアプリケーションやWebサービス(以後、あなたのアプリと呼びます)をMisskeyに登録します。 [デベロッパーセンター](/dev)にアクセスし、「アプリ > アプリ作成」からアプリを作成してください。  登録が済むとあなたのアプリのシークレットキーが入手できます。このシークレットキーは後で使用します。  > アプリに成りすまされる可能性があるため、極力このシークレットキーは公開しないようにしてください。</p>  #### 2.ユーザーに認証させる アプリを使ってもらうには、ユーザーにアカウントへのアクセスの許可をもらう必要があります。  認証セッションを開始するには、[https://misskey.kurume-nct.com/api/auth/session/generate](#operation/auth/session/generate) へパラメータに`appSecret`としてシークレットキーを含めたリクエストを送信します。 レスポンスとして認証セッションのトークンや認証フォームのURLが取得できるので、認証フォームのURLをブラウザで表示し、ユーザーにフォームを提示してください。  あなたのアプリがコールバックURLを設定している場合、 ユーザーがあなたのアプリの連携を許可すると設定しているコールバックURLに`token`という名前でセッションのトークンが含まれたクエリを付けてリダイレクトします。  あなたのアプリがコールバックURLを設定していない場合、ユーザーがあなたのアプリの連携を許可したことを(何らかの方法で(たとえばボタンを押させるなど))確認出来るようにしてください。  #### 3.アクセストークンを取得する ユーザーが連携を許可したら、[https://misskey.kurume-nct.com/api/auth/session/userkey](#operation/auth/session/userkey) へリクエストを送信します。  上手くいけば、認証したユーザーのアクセストークンがレスポンスとして取得できます。おめでとうございます！  アクセストークンが取得できたら、*「ユーザーのアクセストークン+あなたのアプリのシークレットキーをsha256したもの」*をAPIキーとして、APIにリクエストできます。  APIキーの生成方法を擬似コードで表すと次のようになります: ``` js const i = sha256(userToken + secretKey); ``` 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject97 struct for InlineObject97
type InlineObject97 struct {
	// 投稿の公開範囲
	Visibility string `json:"visibility,omitempty"`
	// (投稿の公開範囲が specified の場合)投稿を閲覧できるユーザー
	VisibleUserIds []string `json:"visibleUserIds,omitempty"`
	// 投稿内容
	Text *string `json:"text,omitempty"`
	// コンテンツの警告。このパラメータを指定すると設定したテキストで投稿のコンテンツを隠す事が出来ます。
	Cw *string `json:"cw,omitempty"`
	// モバイルデバイスからの投稿か否か。
	ViaMobile bool `json:"viaMobile,omitempty"`
	// ローカルのみに投稿か否か。
	LocalOnly bool `json:"localOnly,omitempty"`
	// 本文からメンションを展開しないか否か。
	NoExtractMentions bool `json:"noExtractMentions,omitempty"`
	// 本文からハッシュタグを展開しないか否か。
	NoExtractHashtags bool `json:"noExtractHashtags,omitempty"`
	// 本文からカスタム絵文字を展開しないか否か。
	NoExtractEmojis bool `json:"noExtractEmojis,omitempty"`
	Geo *NotesCreateGeo `json:"geo,omitempty"`
	// 添付するファイル
	FileIds []string `json:"fileIds,omitempty"`
	// 添付するファイル (このパラメータは廃止予定です。代わりに fileIds を使ってください。)
	MediaIds []string `json:"mediaIds,omitempty"`
	// 返信対象
	ReplyId string `json:"replyId,omitempty"`
	// Renote対象
	RenoteId string `json:"renoteId,omitempty"`
	Poll NotesCreatePoll `json:"poll,omitempty"`
}
