## How many Go programmers are there in the world?
How many Go programmers are there in the world? Think of a number and hold it in your head, we’ll come back to it at the end of this talk.

## Go プログラマーは世界に何人いるでしょうか？
Go プログラマーは世界に何人いるでしょう？思い浮かんだ数字をトークの最後まで覚えておいてください。

## Code review
Who here does code review as part of their job? [the entire room raised their hand, which was encouraging]. Okay, why do you do code review? [someone shouted out “to stop bad code”]

If code review is there to catch bad code, then how do you know if the code you’re reviewing is good, or bad?

Now it’s fine to say “that code is ugly” or ”wow that source code is beautiful”, just as you might say “this painting is beautiful” or “this room is beautiful” but these are subjective terms, and I’m looking for objective ways to talk about the properties of good or bad code.

## コードレビュー
仕事でコードレビューを行っている人はいますか？
（心強い事に、至る所で手が上がった。）
良いですね。ではなぜコードレビューをするのでしょう？
（誰かが「悪いコードを無くすため！」と叫んだ）

もしコードレビューが悪いコードを見つけ出すためにあるとして、どうやってレビューしているコードを良いか悪いか判断しますか？

もちろん「この絵は美しい」「この部屋は美しい」と言うように「このコードはひどい」または「このコードは美しい」と言う事もできるでしょう。しかしそれは主観的な用語です。
私は客観的な方法で良い、もしくは悪いコードの性質を語る方法を探しています。

## Bad code
What are some of the properties of bad code that you might pick up on in code review?

Rigid. Is the code rigid? Does it have a straight jacket of overbearing types and parameters, that making modification difficult?
Fragile. Is the code fragile? Does the slightest change ripple through the code base causing untold havoc?
Immobile. Is the code hard to refactor? Is it one keystroke away from an import loop?
Complex. Is there code for the sake of having code, are things over-engineered?
Verbose. Is it just exhausting to use the code? When you look at it, can you even tell what this code is trying to do?
Are these positive sounding words? Would you be pleased to see these words used in a review of your code?

Probably not.

## 悪いコードの性質
コードレビューで悪いコードだと思う特徴はなんでしょう？

硬すぎる
ガチガチなコード。変更を加える事が困難な程に型とパラメーターが強引で堅苦しかったりします。

壊れやすい
フラジャイルなコード。
小さな変更が全体に波及し甚大な被害を引き起こします。

移動できない
リファクターが難しいコード。
インポート箇所から遠すぎるコード
FIXME:

複雑すぎ
過剰に設計されたコードのためのコードがある。

くどい
使おうとして疲れてしまうコード。何しているのか説明できないような奴です。

これらの言葉を何か良い意味への言い換えができるでしょうか？
あなたのコードがレビューでこの様な表現をされて嬉しいと思いますか？

おそらく違うでしょう。

