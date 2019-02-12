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

## Good design
But this is an improvement, now we can say things like “I don’t like this because it’s too hard to modify”, or “I don’t like this because i cannot tell what the code is trying to do”, but what about leading with the positive?

Wouldn’t it be great if there were some ways to describe the properties of good design, not just bad design, and to be able to do so in objective terms?

TODO: but what about leading with the positive?

しかしこれは進歩です。これで私たちは「このコードは変更しづらすぎるから好きではない」もしくは「このコードが何をしようとしているか分からないので良いと思わない」といった表現ができるようになります。しかしこれは何を意味するのでしょう？

もし悪いデザインだけではなく、良いデザインの特性を客観的な用語で説明できる様になれば最高じゃないでしょうか？

## SOLID
In 2002 Robert Martin published his book, Agile Software Development, Principles, Patterns, and Practices. In it he described five principles of reusable software design, which he called the SOLID principles, after the first letters in their names.

Single Responsibility Principle
Open / Closed Principle
Liskov Substitution Principle
Interface Segregation Principle
Dependency Inversion Principle
This book is a little dated, the languages that it talks about are the ones in use more than a decade ago. But, perhaps there are some aspects of the SOLID principles that may give us a clue about how to talk about a well designed Go programs.

So this is what I want to spend some time discussing with you this morning.

2002年。ロバート・マーティンはアジャイルソフトウェア開発・原則・実践パターンを発表しました。その中で、再利用可能なソフトウェア設計の5原則を説明し、イニシャルをとってSOLID原則と名付けました。

- 単一責任原則
- オープン・クローズド原則
- リスコフ置換の原則
- インターフェース分離の原則
- 依存関係逆転の原則

この本は少し時代遅れです。語られている言語は10年前に使われていたものです。しかし、おそらく SOLID 原則はよく設計された Go プログラムを説明する方法についての手がかりとなる側面があります。

これがこの朝に私が皆さんにお伝えする時間を頂きたい話しです。

## Single Responsibility Principle
The first principle of SOLID, the S, is the single responsibility principle.

    A class should have one, and only one, reason to change.
    –Robert C Martin

Now Go obviously doesn’t have classes—instead we have the far more powerful notion of composition—but if you can look past the use of the word class, I think there is some value here.

Why is it important that a piece of code should have only one reason for change? Well, as distressing as the idea that your own code may change, it is far more distressing to discover that code your code depends on is changing under your feet. And when your code does have to change, it should do so in response to a direct stimuli, it shouldn’t be a victim of collateral damage.

So code that has a single responsibility therefore has the fewest reasons to change.

TODO: notion
TODO: distressing
TODO: stimuli


SOLID の最初の S は単一責任原則です。

    Class を変更する理由は一つ、ただ一つであるべきだ。

Go には一切の Class 機能がありません。代わりに遥かに強力な Composition という概念があります。しかし Class という言葉の使用に目をつむれば、そこには価値ある内容があります。

1つのコードの変更理由が1つしかない事がなぜ重要なのでしょう？ 自分の書いたコードが変更されるだろう事を考えるのは面倒な事ですが、あなたのコードが依存するコードが見えない場所で変化している事を見つけるのは、それより遥かに面倒な事です。
コードは直接的な要因によって変更されるべきであり、副次的な被害によって行うべきではありません。

単一の責任をもつコードは、したがってもっとも変更する理由がすくないコードになります。

## Coupling & Cohesion
Two words that describe how easy or difficult it is to change a piece of software are coupling and cohesion.

Coupling is simply a word that describes two things changing together–a movement in one induces a movement in another.

A related, but separate, notion is the idea of cohesion, a force of mutual attraction.

In the context of software, cohesion is the property of describing pieces of code are naturally attracted to one another.

To describe the units of coupling and cohesion in a Go program, we might talk about functions and methods, as is very common when discussing SRP but I believe it starts with Go’s package model.

TODO: induces
TODO: mutual
TODO: attraction

結合と凝縮
ソフトウェアを変更するのがどれほど簡単、もしくは難しいかを表すの表現が2つあります。結合と結束です。

結合は簡単な表現です。二つの事が同時に変化する事を表し、一方が動けばもう一方を動かします。

結束はこれに関連はしているが別の、相互に引きよせる力の概念です。

ソフトウェアの文脈では、結束は自然と一方を引き寄せてしまう性質のコードを表現する語彙です。

Go プログラムでの結合と結束の単位を説明するのに、単一責任の原則の解説で一般的な、関数とメソッドの説明を使う事も出来ますが、Goではパッケージモデルの説明から始めるのがよいと考えています。

## Package names
In Go, all code lives inside a package, and a well designed package starts with its name. A package’s name is both a description of its purpose, and a name space prefix. Some examples of good packages from the Go standard library might be:

net/http, which provides http clients and servers.
os/exec, which runs external commands.
encoding/json, which implements encoding and decoding of JSON documents.

When you use another package’s symbols inside your own this is accomplished by the `import` declaration, which establishes a source level coupling between two packages. They now know about each other.

パッケージ名
Go では全てのコードがパッケージに収まっており、よく設計されたパッケージはそのパッケージ名から始まっています。パッケージ名はその用途の説明と、名前空間の接頭辞の両方を担っています。Go 標準ライブラリからよく設計されたパッケージの例を選んでみましょう。

net/http http クライアントとサーバーを提供します。
os/exec 外部コマンドを実行します。
encoding/json json のエンコードと JSON ドキュメントのデコード実装です。

コード中で別のパッケージのシンボルを使う場合、import 定義をすることで2パッケージ間のソースレベルの結合が行われ、利用可能になります。そしてパッケージたちはお互いのことをしっている状態となります。
