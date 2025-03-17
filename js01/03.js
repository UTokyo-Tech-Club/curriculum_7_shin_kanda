const N = Number(window.prompt("自然数を入力してください"));

if (N>0&&Number.isInteger(N)) {
/* 方法1: 和の公式を用いる
  ans1 に N (N + 1) / 2　の計算結果を代入する */
  let ans1 = N * (N + 1) / 2;
/* 方法2: 繰り返しを用いる
  ans2 の初期値を 0 とし, 1 から N を順番に加えてゆく */
  let ans2 = 0;
  for (let i = 1; i <= N; i++) {
    ans2 += i;
  }
/* ２つの結果を比較
  もし等しければ計算結果を, 等しくなければ "Failed to compute" を出力
  三項演算子を使ってみよう
  document.write(
    「もし ans1 == ans2 なら ans2 を そうでなければ "Failed to compute"　を返す」
  )
*/

  document.write(ans1 == ans2 ? ans2 : "Failed to compute");
}



