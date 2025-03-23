'use strict';

const ul=document.getElementById('list');
for (let i=0; i<5; i+=1) {
  const li=document.createElement('li');
  li.textContent=`子要素${i}号`;
  ul.appendChild(li);
}
