<script>
  import {fade} from "svelte/transition";
  import Modal from "../component/Modal.svelte";
  import PanelFull from "../component/Panel.svelte";
  import Placeholder from "../component/Placeholder.svelte";
  let listpasaran = [];
  let listnews = [];
  let listkeluaran = [];
  let listpaito_minggu = [];
  let listpaito_senin = [];
  let listpaito_selasa = [];
  let listpaito_rabu = [];
  let listpaito_kamis = [];
  let listpaito_jumat = [];
  let listpaito_sabtu = [];
  let record_minggu = "";
  let record_senin = "";
  let record_selasa = "";
  let record_rabu = "";
  let record_kamis = "";
  let record_jumat = "";
  let record_sabtu = "";
  let myModal = "";
  async function initNews() {
    const resnews = await fetch("/api/listnews", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({}),
    });
    if (!resnews.ok) {
      const pasarMessage = `An error has occured: ${resnews.status}`;
      throw new Error(pasarMessage);
    } else {
      const jsonnews = await resnews.json();
      if (jsonnews.status == 200) {
        let record = jsonnews.record;
        if (record != null) {
          for (var i = 0; i < record.length; i++) {
            listnews = [
              ...listnews,
              {
                // news_title: temp.replace("undefined", ""),
                news_title: record[i]["news_title"],
                news_descp: record[i]["news_descp"],
                news_url: record[i]["news_url"],
                news_image: record[i]["news_image"],
              },
            ];
          }
        } else {
          alert("Error");
        }
      } else {
        alert("Error");
      }
    }
  }
  async function initPasaran() {
    const respasaran = await fetch("/api/listpasaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({}),
    });
    if (!respasaran.ok) {
      const pasarMessage = `An error has occured: ${resPasar.status}`;
      throw new Error(pasarMessage);
    } else {
      const jsonpasaran = await respasaran.json();
      if (jsonpasaran.status == 200) {
        let record = jsonpasaran.record;
        if (record != null) {
          for (var i = 0; i < record.length; i++) {
            listpasaran = [
              ...listpasaran,
              {
                pasaran_id: record[i]["pasaran_id"],
                pasaran_name: record[i]["pasaran_name"],
                pasaran_url: record[i]["pasaran_url"],
                pasaran_diundi: record[i]["pasaran_diundi"],
                pasaran_jamjadwal: record[i]["pasaran_jamjadwal"],
                pasaran_datekeluaran: record[i]["pasaran_datekeluaran"],
                pasaran_keluaran: record[i]["pasaran_keluaran"],
                pasaran_dateprediksi: record[i]["pasaran_dateprediksi"],
                pasaran_bbfsprediksi: record[i]["pasaran_bbfsprediksi"],
                pasaran_nomorprediksi: record[i]["pasaran_nomorprediksi"],
              },
            ];
          }
        } else {
          alert("Error");
        }
      } else {
        alert("Error");
      }
    }
  }
  async function initKeluaran(e) {
    listkeluaran = [];
    listpaito_minggu = [];
    listpaito_senin = [];
    listpaito_selasa = [];
    listpaito_rabu = [];
    listpaito_kamis = [];
    listpaito_jumat = [];
    listpaito_sabtu = [];
    const respasaran = await fetch("/api/listkeluaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        pasaran: e,
      }),
    });
    if (!respasaran.ok) {
      const pasarMessage = `An error has occured: ${resPasar.status}`;
      throw new Error(pasarMessage);
    } else {
      const jsonpasaran = await respasaran.json();
      console.log(jsonpasaran);
      if (jsonpasaran.status == 200) {
        let record = jsonpasaran.record;
        record_minggu = jsonpasaran.paito_minggu;
        record_senin = jsonpasaran.paito_senin;
        record_selasa = jsonpasaran.paito_selasa;
        record_rabu = jsonpasaran.paito_rabu;
        record_kamis = jsonpasaran.paito_kamis;
        record_jumat = jsonpasaran.paito_jumat;
        record_sabtu = jsonpasaran.paito_sabtu;
        if (record != null) {
          for (var i = 0; i < record.length; i++) {
            listkeluaran = [
              ...listkeluaran,
              {
                keluaran_datekeluaran: record[i]["keluaran_datekeluaran"],
                keluaran_periode: record[i]["keluaran_periode"],
                keluaran_nomor: record[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_minggu != null) {
          for (var i = 0; i < record_minggu.length; i++) {
            listpaito_minggu = [
              ...listpaito_minggu,
              {
                keluaran_nomor: record_minggu[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_senin != null) {
          for (var i = 0; i < record_senin.length; i++) {
            listpaito_senin = [
              ...listpaito_senin,
              {
                keluaran_nomor: record_senin[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_selasa != null) {
          for (var i = 0; i < record_selasa.length; i++) {
            listpaito_selasa = [
              ...listpaito_selasa,
              {
                keluaran_nomor: record_selasa[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_rabu != null) {
          for (var i = 0; i < record_rabu.length; i++) {
            listpaito_rabu = [
              ...listpaito_rabu,
              {
                keluaran_nomor: record_rabu[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_kamis != null) {
          for (var i = 0; i < record_kamis.length; i++) {
            listpaito_kamis = [
              ...listpaito_kamis,
              {
                keluaran_nomor: record_kamis[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_jumat != null) {
          for (var i = 0; i < record_jumat.length; i++) {
            listpaito_jumat = [
              ...listpaito_jumat,
              {
                keluaran_nomor: record_jumat[i]["keluaran_nomor"],
              },
            ];
          }
        }
        if (record_sabtu != null) {
          for (var i = 0; i < record_sabtu.length; i++) {
            listpaito_sabtu = [
              ...listpaito_sabtu,
              {
                keluaran_nomor: record_sabtu[i]["keluaran_nomor"],
              },
            ];
          }
        }
      } else {
        alert("Error");
      }
    }
  }
  const handlecallresult = (e) => {
    myModal = new bootstrap.Modal(document.getElementById("modalresult"));
    myModal.show();
    initKeluaran(e);
  };
  initPasaran();
  initNews();

  let num1 = "a00";
  let num2 = "a00";
  let num3 = "a00";
  let num4 = "a00";
  let num5 = "a00";
  let num6 = "a00";
  let doing = false;

  let spin = [
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
    new Audio("/sounds/spin.mp3"),
  ];
  const generator = () => {
    if (doing) {
      return null;
    }
    doing = true;
    let numChanges = randomInt(1, 4) * 7;
    let numberAngka = numChanges + randomInt(1, 7);
    let numberAngka2 = numChanges + 2 * 7 + randomInt(1, 7);
    let numberAngka3 = numChanges + 4 * 7 + randomInt(1, 7);
    let numberAngka4 = numChanges + 6 * 7 + randomInt(1, 7);
    let numberAngka5 = numChanges + 8 * 7 + randomInt(1, 7);
    let numberAngka6 = numChanges + 8 * 7 + randomInt(1, 7);

    let i1 = 0;
    let i2 = 0;
    let i3 = 0;
    let i4 = 0;
    let i5 = 0;
    let i6 = 0;
    let sound = 0;
    num1 = setInterval(spin1, 50);
    num2 = setInterval(spin2, 50);
    num3 = setInterval(spin3, 50);
    num4 = setInterval(spin4, 50);
    num5 = setInterval(spin5, 50);
    num6 = setInterval(spin6, 50);
    function spin1() {
      i1++;
      if (i1 >= numberAngka) {
        // coin[0].play()
        clearInterval(num1);
        return null;
      }
      num1 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }
    function spin2() {
      i2++;
      if (i2 >= numberAngka2) {
        // coin[1].play()
        clearInterval(num2);
        return null;
      }
      num2 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }
    function spin3() {
      i3++;
      if (i3 >= numberAngka3) {
        // coin[2].play()
        clearInterval(num3);
        return null;
      }
      num3 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }
    function spin4() {
      i4++;
      if (i4 >= numberAngka4) {
        // coin[3].play()
        clearInterval(num4);
        return null;
      }
      num4 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }
    function spin5() {
      i5++;
      if (i5 >= numberAngka5) {
        // coin[4].play()
        clearInterval(num5);
        return null;
      }

      num5 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }
    function spin6() {
      i6++;
      if (i6 >= numberAngka6) {
        // coin[5].play()
        clearInterval(num6);
        testWin();
        return null;
      }
      sound = 0;
      spin[sound].play();
      num6 = "a" + parseInt(chance.integer({ min: 0, max: 9 }));
    }

    num1 = "a00";
    num2 = "a00";
    num3 = "a00";
    num4 = "a00";
    num5 = "a00";
    num6 = "a00";
    doing = false;
  };
  function testWin() {
    doing = false;
  }
  function randomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
  }
</script>

<div class="row" style="padding-top:10px;padding-bottom:10px;">
  <div class="col-sm-6" style="margin:0px;padding:0px 3px 0px 0px">
    {#if listpasaran != ""}
      <div class="card" style="background-color:#2c2c2c;border:none;">
        <div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
          <h1 style="font-size: 16px;color:white;font-weight:bold;">Result Togel</h1>
        </div>
        <div class="card-body" style="margin: 0px;padding:0px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;">
          <table class="table table-sm" style="width: 100%;">
            <thead>
              <tr>
                <th
                  width="*"
                  style="text-align:left;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;" >PASARAN</th>
                <th
                  width="40%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">TANGGAL</th>
                <th
                  width="30%"
                  style="text-align:left;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">HARI</th>
                <th
                  width="20%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">JADWAL</th>
                <th
                  width="20%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">KELUARAN</th>
              </tr>
            </thead>
            <tbody style="border-top:none;border-bottom-color: #2c2c2c;">
              {#each listpasaran as rec}
                <tr in:fade>
                  <td
                    NOWRAP
                    style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">
                    <a
                      href={rec.pasaran_url}
                      target="_blank"
                      style="color:white;">
                      {rec.pasaran_name}
                    </a>
                  </td>
                  <td
                    NOWRAP
                    style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">{rec.pasaran_datekeluaran}</td>
                  <td
                    NOWRAP
                    style="text-align:left;vertical-align:top;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">{rec.pasaran_diundi}</td>
                  <td
                    NOWRAP
                    style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">{rec.pasaran_jamjadwal}</td>
                  <td
                    on:click={() => {
                      handlecallresult(rec.pasaran_id);
                    }}
                    NOWRAP
                    style="text-align:center;vertical-align:top;color:#ffbe00;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;font-weight:bold;cursor:pointer;text-decoration:underline;">
                    {rec.pasaran_keluaran}
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    {:else}
      <Placeholder total_placeholder="3" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
    {/if}
    {#if listpasaran != ""}
      <div class="card" style="background-color:#2c2c2c;border:none;margin-top:5px;">
        <div class="card-header" style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;">
            <h1 style="font-size: 16px;color:white;font-weight:bold;">Prediksi Togel</h1>
        </div>
        <div class="card-body"
          style="margin: 0px;padding:0px;background-color: #2c2c2c;border-bottom:1px solid #191c1f;">
          <table class="table table-sm" style="width: 100%;">
            <thead>
              <tr>
                <th
                  width="*"
                  style="text-align:left;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">PASARAN</th>
                <th
                  width="40%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">TANGGAL</th>
                <th
                  width="20%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">BBFS</th>
                <th
                  width="20%"
                  style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;font-size: 13px;border-bottom:1px solid #2c2c2c;">NOMOR</th>
              </tr>
            </thead>
            <tbody style="border-top:none;border-bottom-color: #2c2c2c;">
              {#each listpasaran as rec}
                <tr in:fade>
                  <td NOWRAP
                    style="text-align:left;vertical-align:top;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">
                    <a
                      href={rec.pasaran_url}
                      target="_blank"
                      style="color:white;">
                      {rec.pasaran_name}
                    </a>
                  </td>
                  <td NOWRAP
                    style="text-align:center;vertical-align:top;color:white;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;">{rec.pasaran_dateprediksi}</td>
                  <td NOWRAP
                    style="text-align:center;vertical-align:top;color:#ffbe00;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;font-weight:bold;">{rec.pasaran_bbfsprediksi}</td
                  >
                  <td NOWRAP
                    style="text-align:center;vertical-align:top;color:#ffbe00;background-color:#2c2c2c;border-bottom:1px solid #2c2c2c;font-size:12px;font-weight:bold;">{rec.pasaran_nomorprediksi}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    {:else}
      <Placeholder total_placeholder="3" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
    {/if}
    <ul class="nav nav-pills">
      <li class="nav-item">
        <button
          class="nav-link active"
          id="pills-angkamain-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-angkamain"
          type="button"
          role="tab"
          aria-controls="pills-angkamain"
          aria-selected="true">Angka Main</button>
      </li>
      <li class="nav-item">
        <button
          class="nav-link"
          id="pills-generator-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-generator"
          type="button"
          role="tab"
          aria-controls="pills-generator"
          aria-selected="false">Generator</button>
      </li>
      <li class="nav-item">
        <button
          class="nav-link"
          id="pills-shio-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-shio"
          type="button"
          role="tab"
          aria-controls="pills-shio"
          aria-selected="false">Table Shio</button>
      </li>
      <li class="nav-item">
        <button
          class="nav-link"
          id="pills-teysen-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-teysen"
          type="button"
          role="tab"
          aria-controls="pills-teysen"
          aria-selected="false">Table Tyesen</button>
      </li>
    </ul>
    <div class="tab-content" id="nav-tabContent">
      <div class="tab-pane fade show active"
        id="pills-angkamain"
        role="tabpanel"
        aria-labelledby="pills-angkamain-tab">
        <div class="card" style="background-color:#ffbe00;border:none;margin-top:5px;">
          <div class="card-body" style="margin: 0px;padding:0px;background-color: #ffbe00;">
            <table class="table table-sm" style="width: 100%;">
              <thead>
                <tr style="background-color: #ffbe00;border-style: none;border-bottom-color: #ffbe00;">
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">HARI</th>
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">PON</th>
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">WAGE</th>
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">KLIWON</th>
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">LEGI</th>
                  <th style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;">PAHING</th>
                </tr>
              </thead>
              <tbody style="border-top:none;border-bottom-color: #e91e65;">
                <tr style="background-color: #ffbe00;border-style: none;border-bottom-color: #ffbe00;">
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>SENIN</b></td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">89102</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">67890</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">23456</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">01234</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">67890</td>
                </tr>
                <tr>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>RABU</b></td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">90123</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">89012</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">67890</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">56789</td>
                  <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">34567</td>
                </tr>
                <tr>
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                    ><b>KAMIS</b></td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >67890</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >23456</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >45678</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >01234</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >23456</td
                  >
                </tr>
                <tr>
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                    ><b>SABTU</b></td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >34567</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >67890</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >23456</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >56789</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >12345</td
                  >
                </tr>
                <tr>
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                    ><b>MINGGU</b></td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >56789</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >01234</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >01234</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >56789</td
                  >
                  <td
                    style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                    >56789</td
                  >
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div class="tab-pane fade"
        id="pills-generator"
        role="tabpanel"
        aria-labelledby="pills-generator-tab">
        <div class="card" style="background-color:#ffbe00;border:none;margin-top:5px;">
          <div class="card-body" style="margin: 0px;padding:0px;background-color: #ffbe00;">
            <center style="margin: 10px 0px 5px 0px;">
              <div id="num1" class={num1} />
              <div id="num2" class={num2} />
              <div id="num3" class={num3} />
              <div id="num4" class={num4} />
              <div id="num5" class={num5} />
              <div id="num6" class={num6} />
              <br />
              <br />
              <button
                on:click={() => {
                  generator();
                }}
                type="button"
                class="btn btn-dark btn-sm">Generate</button
              >
            </center>
          </div>
        </div>
      </div>
      <div class="tab-pane fade"
        id="pills-shio"
        role="tabpanel"
        aria-labelledby="pills-shio-tab">
        <div class="card" style="background-color:#ffbe00;border:none;margin-top:5px;">
            <div class="card-body" style="margin: 0px;padding:0px;background-color: #ffbe00;">
              <div class="table-responsive">
                <table class="table table-sm" style="width: 100%;">
                  <thead>
                    <tr>
                      <th
                        style="text-align:left;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;border-bottom:1px solid #ffbe00;">SHIO</th>
                      <th
                        colspan="9"
                        style="text-align:center;vertical-align:top;font-size:13px;color:black;background-color:#ffbe00;border-bottom:1px solid #ffbe00;">ANGKA</th>
                    </tr>
                  </thead>
                  <tbody style="border-top:none;border-bottom-color: #e91e65;">
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>KERBAU</b></td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">01</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">13</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">25</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">37</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">49</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">61</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">73</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">85</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">97</td>
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>TIKUS</b></td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">02</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">14</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">26</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">38</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">50</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">62</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">74</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">86</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">98</td>
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>BABI</b></td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">03</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">15</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">28</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">39</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">51</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">63</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">75</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">86</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">99</td>
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>ANJING</b></td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">04</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">16</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >29</td>
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >40</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >52</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >64</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >75</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >87</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >00</td
                      >
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>AYAM</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >05</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >17</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >30</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >41</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >53</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >65</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >76</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >88</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>MONYET</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >06</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >18</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >31</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >42</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >54</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >66</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >79</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >89</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>KAMBING</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >06</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >19</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >32</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >43</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >55</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >67</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >80</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >90</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>KUDA</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >07</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >20</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >33</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >44</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >56</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >68</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >81</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >91</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>ULAR</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >08</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >21</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >34</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >45</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >57</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >69</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >82</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >92</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>NAGA</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >09</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >22</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >35</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >46</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >58</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >70</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >83</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >93</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>KELINCI</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >10</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >23</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >36</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >47</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >59</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >71</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >84</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >94</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                      />
                    </tr>
                    <tr>
                      <td
                        style="text-align:left;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"
                        ><b>MACAN</b></td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >11</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >24</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >37</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >48</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >60</td
                      >
                      <td
                        style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"
                        >72</td
                      >
                      <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">85</td>
                      <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">95</td>
                      <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"/>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
        </div>
      </div>
      <div class="tab-pane fade"
        id="pills-teysen"
        role="tabpanel"
        aria-labelledby="pills-teysen-tab">
        <div class="card"
          style="background-color:#ffbe00;border:none;margin-top:5px;">
          <div class="card-body" style="margin: 0px;padding:0px;background-color: #ffbe00;">
            <div class="table-responsive">
              <table class="table table-sm" style="width: 100%;">
                <tbody>
                  <tr>
                    <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;" nowrap>
                      01 05 95 12 45<br />
                      02 16 53 09 35<br />
                      03 32 52 85 52<br />
                      04 12 65 05 15<br />
                      05 01 89 10 39<br />
                      06 20 91 51 41<br />
                      07 24 58 57 08<br />
                      08 17 57 04 07<br />
                      09 33 87 88 37<br />
                      10 18 82 03 32<br />
                      11 15 77 02 27<br />
                      12 04 69 17 19<br />
                      13 14 79 07 29<br />
                      14 13 96 08 46<br />
                      15 11 54 00 04<br />
                      16 02 74 15 24<br />
                      17 08 88 13 38<br />
                      18 10 78 01 28<br />
                      19 27 62 54 12<br />
                      20 06 72 19 22
                    </td>
                    <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;" nowrap>
                      21 22 93 55 43<br />
                      22 21 70 50 20<br />
                      23 30 84 81 34<br />
                      24 07 66 14 16<br />
                      25 35 85 82 03<br />
                      26 31 90 80 40<br />
                      27 19 61 06 11<br />
                      28 29 69 56 18<br />
                      29 28 63 53 13<br />
                      30 23 99 58 49<br />
                      31 26 94 59 44<br />
                      32 03 60 18 10<br />
                      33 09 86 16 36<br />
                      34 36 73 89 23<br />
                      35 25 75 52 02<br />
                      36 34 83 87 33<br />
                      37 38 59 83 09<br />
                      38 37 67 84 17<br />
                      39 44 55 77 05<br />
                      40 43 76 78 26
                    </td>
                    <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;" nowrap>
                      41 49 56 76 06<br />
                      42 45 97 72 47<br />
                      43 40 71 41 21<br />
                      44 39 81 86 31<br />
                      45 42 51 75 01<br />
                      46 48 64 73 14<br />
                      47 50 92 21 42<br />
                      48 46 00 79 50<br />
                      49 41 80 70 30<br />
                      50 47 98 74 48<br />
                      51 55 45 22 95<br />
                      52 66 03 99 85<br />
                      53 82 02 35 75<br />
                      54 62 15 95 65<br />
                      55 51 39 20 89<br />
                      56 70 41 71 91<br />
                      57 74 08 47 58<br />
                      58 67 07 94 57<br />
                      59 83 37 38 87<br />
                      60 68 32 93 82
                    </td>
                    <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;" nowrap>
                      61 65 27 92 77<br />
                      62 54 19 27 69<br />
                      63 64 29 97 79<br />
                      64 63 46 98 96<br />
                      65 61 04 90 54<br />
                      66 52 24 25 74<br />
                      67 58 38 23 88<br />
                      68 60 28 91 78<br />
                      69 77 12 44 62<br />
                      70 56 22 29 72<br />
                      71 72 43 45 93<br />
                      72 71 20 40 70<br />
                      73 80 34 31 84<br />
                      74 57 16 24 66<br />
                      75 85 35 32 53<br />
                      76 81 40 30 90<br />
                      77 69 11 96 61<br />
                      78 79 18 46 68<br />
                      79 78 13 43 83<br />
                      80 73 49 48 99
                    </td>
                    <td style="text-align:center;vertical-align:top;font-size:12px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;" nowrap>
                      81 76 44 49 94<br />
                      82 53 10 28 60<br />
                      83 59 36 26 86<br />
                      84 86 23 39 73<br />
                      85 75 26 42 52<br />
                      86 84 33 37 83<br />
                      87 88 09 33 59<br />
                      88 87 17 34 64<br />
                      89 94 05 67 55<br />
                      90 93 26 68 76<br />
                      91 99 06 66 56<br />
                      92 95 47 62 97<br />
                      93 90 21 61 71<br />
                      94 89 31 36 81<br />
                      95 92 01 65 51<br />
                      96 98 14 63 64<br />
                      97 00 42 11 92<br />
                      98 86 50 69 00<br />
                      99 91 30 60 80<br />
                      00 97 48 64 98
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="col-sm-6" style="margin:0px;padding:0px 0px 0px 3px;">
    {#if listnews != ""}
    <PanelFull
        header={true}
        footer={false}
        card_style="background-color:#2c2c2c;border:none;padding:0px;margin:0px;"
        header_style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2c2c2c;border-bottom:2px solid #ed247a;"
        body_style="margin: 0px 0px 30px 0px;padding:0px;background-color: #2c2c2c;border-bottom:1px solid #2c2c2c;height:765px;">
        <slot:template slot="header">
          <h1 style="font-size: 16px;color:white;font-weight:bold;">News</h1>
        </slot:template>
        <slot:template slot="body">
          {#each listnews as rec}
            <a in:fade href="{rec.news_url}" style="color:white;text-decoration: none;" target="_blank" alt="{rec.news_title}">
              <div class="card"
                style="background-color:#2c2c2c;border:none;margin:5px;border-bottom:1px solid #ed247a;">
                <img src={rec.news_image} class="card-img-top" alt="{rec.news_title}" />
                <div class="card-body" style="background-color:none;border:none;padding:5px 0px 0px 0px;margin:0px;">
                  <h1 class="card-title" style="color:white;font-size:15px;text-decoration: underline;">{rec.news_title}</h1>
                  <p style="font-size:13px;">
                    {rec.news_descp}
                  </p>
                </div>
              </div>
            </a>
          {/each}
        </slot:template>
    </PanelFull>
    {:else}
      <Placeholder total_placeholder="10" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
    {/if}
  </div>
</div>
<Modal
  modal_id="modalresult"
  modal_size="modal-dialog-centered"
  modal_title="LAST RESULT"
  modal_modal_css="background-color:#191c1f;border:none;"
  modal_header_css="color:white;font-weight:bold;"
  modal_body_css="height:500px;padding:5px;"
  modal_footer_css="padding:5px;"
  modal_footer={false}
  header_flag={true}>
  <slot:template slot="header">
    <ul class="nav nav-pills">
      <li class="nav-item">
        <button
          class="nav-link active"
          id="pills-all-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-all"
          type="button"
          role="tab"
          aria-controls="pills-all"
          aria-selected="true">RESULT</button
        >
      </li>
      <li class="nav-item">
        <button
          class="nav-link"
          id="pills-paito-tab"
          data-bs-toggle="pill"
          data-bs-target="#pills-paito"
          type="button"
          role="tab"
          aria-controls="pills-paito"
          aria-selected="false">PAITO</button
        >
      </li>
    </ul>
  </slot:template>
  <slot:template slot="body">
    <div class="tab-content" id="nav-tabContent">
      <div class="tab-pane fade show active"
        id="pills-all"
        role="tabpanel"
        aria-labelledby="pills-all-tab">
        {#if listkeluaran != ""}
          <table class="table table-striped table-sm">
            <thead>
              <tr style="background-color: #191c1f;border-bottom-color: #e80650;">
                <th width="*"
                  style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">DATE</th>
                <th width="40%"
                  style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">PERIODE</th>
                <th width="40%"
                  style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">NOMOR</th>
              </tr>
            </thead>
            <tbody style="border-top:none;border-bottom-color: #2c2c2c;">
              {#each listkeluaran as rec}
                <tr in:fade style="">
                  <td NOWRAP
                    style="text-align: center;vertical-align:top;font-size: 13px;color:#ffffff;border-bottom:1px solid #2c2c2c;">{rec.keluaran_datekeluaran}</td>
                  <td NOWRAP
                    style="text-align: center;vertical-align:top;font-size: 13px;color:#ffffff;border-bottom:1px solid #2c2c2c;">{rec.keluaran_periode}</td>
                  <td NOWRAP
                    style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;border-bottom:1px solid #2c2c2c;">{rec.keluaran_nomor}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        {:else}
          <Placeholder total_placeholder="5" card_style="background-color:#2c2c2c;border:none;margin-top:5px;" />
        {/if}
      </div>
      <div class="tab-pane fade"
        id="pills-paito"
        role="tabpanel"
        aria-labelledby="pills-paito-tab">
        <table class="table table-striped table-sm">
          <thead>
            <tr style="background-color: #191c1f;border-bottom-color: #e80650;">
              <th width="*"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">MINGGU</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">SENIN</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">SELASA</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">RABU</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">KAMIS</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">JUMAT</th>
              <th width="25%"
                style="text-align: center;vertical-align:top;font-size: 14px;color:#ffffff;">SABTU</th>
            </tr>
          </thead>
          <tbody style="border-top:none;border-bottom-color: #2c2c2c;">
            <tr style="">
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_minggu as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_senin as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_selasa as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_rabu as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_kamis as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_jumat as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
              <td style="text-align: center;vertical-align:top;font-size: 13px;color:#ffbe00;">
                {#each listpaito_sabtu as rec}
                  {rec.keluaran_nomor} <br />
                {/each}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </slot:template>
</Modal>

<style>
  #num1,
  #num2,
  #num3,
  #num4,
  #num5,
  #num6 {
    display: inline-block;
    margin-top: 2px;
    margin-left: 2px;
    margin-right: 2px;
    background-size: 70px;
    width: 70px;
    height: 70px;
  }
  .a00 {
    background-image: url(number/ball-null.svg);
  }
  .a0 {
    background-image: url(number/ball-0.svg);
  }
  .a1 {
    background-image: url(number/ball-1.svg);
  }
  .a2 {
    background-image: url(number/ball-2.svg);
  }
  .a3 {
    background-image: url(number/ball-3.svg);
  }
  .a4 {
    background-image: url(number/ball-4.svg);
  }
  .a5 {
    background-image: url(number/ball-5.svg);
  }
  .a6 {
    background-image: url(number/ball-6.svg);
  }
  .a7 {
    background-image: url(number/ball-7.svg);
  }
  .a8 {
    background-image: url(number/ball-8.svg);
  }
  .a9 {
    background-image: url(number/ball-9.svg);
  }
  .nav-link {
      display: block;
      padding: .5rem 1rem;
      color: white;
      text-decoration: none;
      transition: color .15s ease-in-out,background-color .15s ease-in-out,border-color .15s ease-in-out;
  }
  .nav-pills .nav-link.active,
  .nav-pills .show > .nav-link {
    color: black;
    background-color: #ffbe00;
  }
  .nav-pills .nav-link {
    border-radius: 0.1rem;
  }
  button {
      margin: 0;
      font-family: inherit;
      font-size: 14px;
      line-height: inherit;
  }
 
</style>
