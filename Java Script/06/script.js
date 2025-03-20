'use strict';

const table=document.querySelector('table');
const tr=document.createElement('tr');
const header=['名前','タイプ','体重','特技'];

for (const h of header){
  const th=document.createElement('th');
  th.textContent=h;
  tr.appendChild(th);
}
table.appendChild(tr);

const evList = [
    { name: 'サンダース', type: 'でんき', weight: 24.5, ability: 'ちくでん' },
    { name: 'ブースター', type: 'ほのお', weight: 25.0, ability: 'もらいび' },
    { name: 'エーフィ', type: 'エスパー', weight: 26.5, ability: 'シンクロ' },
    { name: 'ブラッキー', type: 'あく', weight: 27.0, ability: 'シンクロ' },
    { name: 'リーフィア', type: 'くさ', weight: 25.5, ability: 'リーフガード' },
    { name: 'ニンフィア', type: 'フェアリー', weight: 23.5, ability: 'メロメロボディ' },
    { name: 'グレイシア', type: 'こおり', weight: 25.9, ability: 'ゆきがくれ' },
    { name: 'シャワーズ', type: 'みず', weight: 29.0, ability: 'ちょすい' },
  ];

for (const ev of evList){
    const tr=document.createElement('tr');
    for (const prop in ev){
        const td=document.createElement('td');
        td.textContent=ev[prop];
        if (prop=='weight'){
            td.textContent+='kg';
        }
        tr.appendChild(td);
    }
    table.appendChild(tr);
}



