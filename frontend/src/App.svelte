<script>
	import Navbar from "./component/Navbar.svelte"

	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}
	console.log(client_device)
</script>


<div class="contentmodif" style="margin-top: -10px;">
	<div class="clearfix"></div><br>
	<div class="content container" style="margin-top: 70px;">
		<div class="row" style="padding-top:10px;padding-bottom:10px;">
			<div class="col-sm-6" style="">
				<div class="card" style="background-color:#191c1f;border:none;">
					<div class="card-header"
						style="padding: 10px 0px 5px 10px;margin:0px;background-color:#191c1f;border-bottom:1px solid #191c1f;">
						<h1 style="font-size: 16px;color:white;font-weight:bold;">Result Togel</h1>
					</div>
					<div
						class="card-body"
						style="margin: 0px;padding:0px;background-color: #191c1f;border-bottom:1px solid #191c1f;">
						<table class="table" style="width: 100%;border-collapse:unset;border-bottom:1px solid #191c1f;" >
							<thead>
								<tr>
									<th style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;font-size: 14px;" width="20%">PASARAN</th>
									<th style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;font-size: 14px;" width="20%">TANGGAL</th>
									<th style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;font-size: 14px;" width="*">KELUARAN&nbsp;</th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#191c1f;border-bottom:1px solid #191c1f;font-size:14px;">1234</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
				<br>
				<div class="card" style="background-color:#cd2a12;border:none;">
					<div class="card-header"
						style="padding: 10px 0px 5px 10px;margin:0px;background-color:#cd2a12;border-bottom:1px solid #cd2a12;">
						<h1 style="font-size: 16px;color:white;font-weight:bold;">Prediksi Togel</h1>
					</div>
					<div
						class="card-body"
						style="margin: 0px;padding:0px;background-color: #cd2a12;border-bottom:1px solid #cd2a12;">
						<table class="table" style="width: 100%;border-collapse:unset;border-bottom:1px solid #cd2a12;" >
							<thead>
								<tr>
									<th style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;font-size: 14px;" width="20%">PASARAN</th>
									<th style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;font-size: 14px;" width="20%">TANGGAL</th>
									<th style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;font-size: 14px;" width="*">BBFS</th>
									<th style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;font-size: 14px;" width="*">NOMOR</th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
								</tr>
								<tr>
									<td style="text-align:left;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">HONGKONG</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">2021-10-10</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
									<td style="text-align:center;vertical-align:top;font-size:12px;color:white;background-color:#cd2a12;border-bottom:1px solid #cd2a12;font-size:14px;">1234</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
				<br>
			</div>
			<div class="col-sm-6" style="">
				<div class="card" style="background-color:#2a2e32;border:none;">
					<div class="card-header"
						style="padding: 10px 0px 5px 10px;margin:0px;background-color:#2a2e32;border-bottom:1px solid #2a2e32;">
						<h1 style="font-size: 16px;color:white;font-weight:bold;">Angka Main Berdasarkan Hari DAN Pasaran</h1>
					</div>
					<div
						class="card-body"
						style="margin: 0px;padding:0px;background-color: #2a2e32;">
						<table class="table" style="width: 100%;border-collapse:unset;" >
							<thead>
								<tr style="background-color: #2a2e32;border-style: none;border-bottom-color: #2a2e32;">
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">HARI</th>
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">PON</th>
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">WAGE</th>
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">KLIWON</th>
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">LEGI</th>
									<th style="text-align:center;vertical-align:top;font-size:15px;color:white;background-color:#2a2e32;">PAHING</th>
								</tr>
							</thead>
							<tbody>
								<tr style="background-color: #2a2e32;border-style: none;border-bottom-color: #2a2e32;">
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;color:white;border-bottom:1px solid #2a2e32;"><b>SENIN</b></td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">89102</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">67890</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">23456</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">01234</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">67890</td>
								</tr>
								<tr>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;color:white;border-bottom:1px solid #2a2e32;"><b>RABU</b></td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">90123</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">89012</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">67890</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">56789</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">34567</td>
								</tr>
								<tr>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;color:white;border-bottom:1px solid #2a2e32;"><b>KAMIS</b></td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">67890</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">23456</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">45678</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">01234</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">23456</td>
								</tr>
								<tr>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;color:white;border-bottom:1px solid #2a2e32;"><b>SABTU</b></td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">34567</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">67890</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">23456</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">56789</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">12345</td>
								</tr>
								<tr>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;color:white;border-bottom:1px solid #2a2e32;"><b>MINGGU</b></td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">56789</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">01234</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">01234</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">56789</td>
									<td style="text-align:center;vertical-align:top;font-size:14px;background:#2a2e32;border-bottom:1px solid #2a2e32;color:white;">56789</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
				
				<div class="clearfix"></div><br>
				<div class="card" style="background-color:#ffbe00;border:none;">
					<div class="card-header"
						style="padding: 10px 0px 5px 10px;margin:0px;background-color:#ffbe00;border-bottom:1px solid #ffbe00;">
						<h1 style="font-size: 16px;color:black;font-weight:bold;">Tabke Shio 2021</h1>
					</div>
					<div
						class="card-body"
						style="margin: 0px;padding:0px;background-color: #ffbe00;">
						<div class="table-responsive">
							<table class="table" style="width: 100%;border-collapse:unset;" >
								<thead>
									<tr>
										<th style="text-align:center;vertical-align:top;font-size:15px;color:black;background-color:#ffbe00;border-bottom:1px solid #ffbe00;">SHIO</th>
										<th colspan="9" style="text-align:center;vertical-align:top;font-size:15px;color:black;background-color:#ffbe00;border-bottom:1px solid #ffbe00;">ANGKA</th>
									</tr>
								</thead>
								<tbody>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>KERBAU</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">01</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">13</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">25</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">37</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">49</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">61</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">73</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">85</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">97</td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>TIKUS</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">02</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">14</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">26</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">38</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">50</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">62</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">74</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">86</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">98</td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>BABI</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">03</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">15</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">28</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">39</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">51</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">63</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">75</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">86</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">99</td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>ANJING</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">04</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">16</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">29</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">40</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">52</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">64</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">75</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">87</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">00</td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>AYAM</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">05</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">17</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">30</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">41</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">53</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">65</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">76</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">88</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>MONYET</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">06</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">18</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">31</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">42</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">54</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">66</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">79</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">89</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>KAMBING</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">06</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">19</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">32</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">43</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">55</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">67</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">80</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">90</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>KUDA</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">07</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">20</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">33</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">44</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">56</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">68</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">81</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">91</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>ULAR</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">08</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">21</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">34</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">45</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">57</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">69</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">82</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">92</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>NAGA</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">09</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">22</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">35</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">46</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">58</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">70</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">83</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">93</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>KELINCI</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">10</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">23</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">36</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">47</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">59</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">71</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">84</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">94</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;color:black;border-bottom:1px solid #ffbe00;"><b>MACAN</b></td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">11</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">24</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">37</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">48</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">60</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">72</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">85</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;">95</td>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#ffbe00;border-bottom:1px solid #ffbe00;color:black;"></td>
									</tr>
								</tbody>
							</table>
						</div>
					</div>
				</div>
				<div class="clearfix"></div><br>
				<div class="card" style="background-color:#ffbe00;border:none;">
					<div class="card-header"
						style="padding: 10px 0px 5px 10px;margin:0px;background-color:#1b9da7;border-bottom:1px solid #1b9da7;">
						<h1 style="font-size: 16px;color:black;font-weight:bold;">Table Tyesen</h1>
					</div>
					<div
						class="card-body"
						style="margin: 0px;padding:0px;background-color: #1b9da7;">
						<div class="table-responsive">
							<table class="table" style="width: 100%;border-collapse:unset;" >
								<tbody>
									<tr>
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#1b9da7;color:black;border-bottom:1px solid #1b9da7;" nowrap>
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
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#1b9da7;color:black;border-bottom:1px solid #1b9da7;" nowrap>
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
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#1b9da7;color:black;border-bottom:1px solid #1b9da7;" nowrap>
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
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#1b9da7;color:black;border-bottom:1px solid #1b9da7;" nowrap>
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
										<td style="text-align:center;vertical-align:top;font-size:14px;background:#1b9da7;color:black;border-bottom:1px solid #1b9da7;" nowrap>
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
	<div class="clearfix"></div><br>
</div>
<style>
	.container {
		width: 1000px;
		max-width: none !important;
	}
	.contentmodif{
    	background-image: url('https://image.winudf.com/v2/image1/Y29tLmFuZHJvbW8uZGV2NzQxODg3LmFwcDk2NDI0M19zY3JlZW5fMF8xNTU4ODMwMDk4XzAwMg/screen-0.jpg?fakeurl=1&type=.jpg');
        background-position:center center;
        background-repeat:no-repeat;
		background-size: cover;
	}
</style>