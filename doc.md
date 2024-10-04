# 設計書
## 概要
jsonで定義したdirectory構造に当てはまっているかどうかを評価する.
  
(npmで提供されていたが、npmを入れたくないので作る)

## 使い方
以下のjson型のobjectでfileとdirectoryを表現する
```struct(不採用)
[
    {
        "name": "name",
        "type": "dir or file",
        "contents": [{...}, ...] // only dir
    }
]
```

```map
{
    "file": null,
    "dir": {...}
}
```

## 実装
tree構造である. bfs(幅優先探索)アルゴリズムを使用する.
lsで表示されたlistをjsonと比較し、directoryがあったら後に同じように行う. 
### flow
- json読み込み. dataに代入
- dataとpathをqueueに入れる
- check_list(queue.pop())
- queueが空になるまで繰り返す
- すべてチェックし終わったら、正常値を返して終了. 

#### キューから受け取ったdirの処理(func check_list) (map)
- lsでfileとdirectoryを表示
- lsで表示した内容をlistにする
- listとdataの数が一致. 一つずつdata["name"]で確認.
- directoryの場合はdata["dirname"]とdirectoryのpathをキューに渡す. 
