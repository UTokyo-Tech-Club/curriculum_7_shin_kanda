const fibo = (N) => {
    if (N === 0) return [0];
    if (N === 1) return [0, 1];

    let f = [0, 1];  // 第0項と第1項を初期値として配列に格納
    for (let i=2; i<=N; i++) {
        // f の最後（n番目）に f[n-1] + f[n-2] を追加する
        f.push(f[i-1] + f[i-2]);
    }
    return f  // 第0項から第N項までが順に格納された配列
};

