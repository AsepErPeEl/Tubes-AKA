async function hitungManual() {
  const a = document.getElementById("a").value;
  const b = document.getElementById("b").value;

  const res = await fetch(`/api/manual?a=${a}&b=${b}`);
  const data = await res.json();

  document.getElementById("manualResult").textContent =
    `a = ${data.a}\n` +
    `b = ${data.b}\n` +
    `Iteratif = ${data.iter_ms.toFixed(6)} ms\n` +
    `Rekursif = ${data.rec_ms.toFixed(6)} ms`;

  // TAMBAH KE TABEL MANUAL
  const tbody = document.querySelector("#manualTable tbody");
  const row = document.createElement("tr");
  row.innerHTML = `
    <td>${data.a}</td>
    <td>${data.b}</td>
    <td>${data.iter_ms.toFixed(6)}</td>
    <td>${data.rec_ms.toFixed(6)}</td>
  `;
  tbody.appendChild(row);
}

async function hitungAuto(n) {
  const res = await fetch(`/api/auto?n=${n}`);
  const data = await res.json();

  const row = document.querySelector(`tr[data-n="${n}"]`);
  row.children[1].textContent = data.iter_ms.toFixed(6);
  row.children[2].textContent = data.rec_ms.toFixed(6);
}
