package main

import (
	"testing"

	"github.com/ikeikeikeike/godic/modules/funcmaps"
)

func TestAutoLink(t *testing.T) {
	names := []string{"様々な試練", "試練", "テスト"}
	html := `
<h1>テスト</h1>
<p><a href="http://example.com/" rel="nofollow"><img src="http://example.com/test.png"></a></p>
<p><code>テスト</code>とは、様々な試練である。</p>
<p>概要</p>
<p><code>テスト</code>の<strong>概要</strong>を<a href="/d/%E6%9B%B8%E3%81%84%E3%81%A6" rel="nofollow">書いて</a>ください。</p>
<p>関連記事</p>
<p><code>テスト</code>に関する<strong>example.orgの記事</strong>を紹介してください。</p>
<p>関連エロ本</p>
<p><code>テスト</code>に関する<strong>example.orgの本</strong>を紹介してください。</p>
<p>関連エロ動画</p>
<p><code>テスト</code>に関する<strong>example.orgの動画</strong>を紹介してください。</p>
<p>関連項目</p>
<p><code>テスト</code>に関する<strong>項目</strong>を紹介してください。</p>
<ul>
<li>項目１</li>
<li>項目２</li>
<li>項目３</li>
</ul>
`
	rHtml := funcmaps.AutoLink(html, names)
	if html == rHtml {
		t.Fatalf("not replace html")
	}
}
