<script>
	import Navbar from "./component/Navbar.svelte"
	import Home from "./pages/Home.svelte"
	import Homemobile from "./pages/Homemobile.svelte"

	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}
	console.log(client_device)
</script>

{#if client_device =="WEBSITE"}
<Navbar />
<div class="contentmodif" style="margin-top: -10px;">
	<div class="clearfix"></div><br>
	<div class="content container" style="margin-top: 30px;">
		<Home />
	</div>
	<div class="clearfix"></div><br>
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
{/if}