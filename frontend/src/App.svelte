<script>
	import Navbar from "./component/Navbar.svelte"
	import Footer from "./component/Footer.svelte"
	import Home from "./pages/Home.svelte"
	import Homemobile from "./pages/Homemobile.svelte"

	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}

	var shios = [
    "tikus",
    "babi",
    "anjing",
    "ayam",
    "monyet",
    "kambing",
    "kuda",
    "ular",
    "naga",
    "kelinci",
    "macan",
    "kerbau",
  ];
  var datashio = [
    ["01", "13", "25", "37", "49", "61", "73", "85", "97"],
    ["02", "14", "26", "38", "50", "62", "74", "86", "98"],
    ["03", "15", "27", "39", "51", "63", "75", "87", "99"],
    ["04", "16", "28", "40", "52", "64", "76", "88", "00"],
    ["05", "17", "29", "41", "53", "65", "77", "89", ""],
    ["06", "18", "30", "42", "54", "66", "78", "90", ""],
    ["07", "19", "31", "43", "55", "67", "79", "91", ""],
    ["08", "20", "32", "44", "56", "68", "80", "92", ""],
    ["09", "21", "33", "45", "57", "69", "81", "93", ""],
    ["10", "22", "34", "46", "58", "70", "82", "94", ""],
    ["11", "23", "35", "47", "59", "71", "83", "95", ""],
    ["12", "24", "36", "48", "60", "72", "84", "96", ""],
  ];

  function rotateRight(arr, count) {
    for (var i = 0; i < count; i++) {
      let last = arr.pop();
      arr.unshift(last);
    }

    return arr;
  }

  function initShio() {
    var year = 2020;
    var getDiffYear = Number(new Date().getFullYear()) - year;
    const res = rotateRight(shios, getDiffYear);
    for (var j = 0; j < datashio.length; j++) {
      datashio[j].unshift(res[j].toUpperCase());
    }
  }

  	initShio();
	console.log(client_device)
</script>

{#if client_device =="WEBSITE"}
<Navbar />
<div class="contentmodif" style="margin-top: -10px;">
	<div class="clearfix"></div><br>
	<div class="content container" style="margin-top: 30px;">
		<Home {datashio} />
	</div>
	<div class="clearfix"></div><br>
	<Footer />
</div>
<style>
	.container {
		width: 1000px;
		max-width: none !important;
	}
</style>
{:else}
<Navbar />
	<div class="container" style="margin-top: 60px;">
		<Homemobile />
	</div>
	<Footer />
{/if}