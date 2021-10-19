<script>
	import { Row, Col, Container } from "sveltestrap";
	
	let token = localStorage.getItem("token");
	let css_loader = "display: none;";
    let msgloader = "";
	let listwinlose = [];
	let pasaran = [];

	async function initapp() {
		const res = await fetch("/api/home", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				page: "DASHBOARD-VIEW",
			}),
		});
		const json = await res.json();
		if (json.status === 400) {
			logout();
		} else {
			initLoad();
		}
	}
	async function initLoad() {
		css_loader = "display: inline-block;";
        msgloader = "Fetch...";
		const res = await fetch("/api/dashboardwinlose", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
			}),
		});
		const json = await res.json();
		let record = json.record;
		if (json.status == 400) {
			// logout();
		} else {
			for (let i = 0; i < record.length; i++) {
				listwinlose = [
					...listwinlose,
					{
						pasaran_name: record[i]["pasaran_name"],
						pasaran_winlose: record[i]["pasaran_detial"],
					},
				];
			}

			for (let j = 0; j < listwinlose.length; j++) {
				let totaldata = listwinlose[j].pasaran_winlose.length;
				let temp = [];
				for (let x = 0; x < totaldata; x++) {
					temp.push(
						parseInt(
							listwinlose[j].pasaran_winlose[x]["Pasaranwinlose"]
						)/1000
					);
				}
				pasaran = [
					...pasaran,
					{
						name: listwinlose[j].pasaran_name,
						data: temp,
					},
				];
			}
			createChart()
			setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
		}
	}
	async function logout() {
		localStorage.clear();
		window.location.href = "/";
	}
	initapp();

	async function createChart() {
		Highcharts.chart("container", {
			chart: {
				type: "column",
			},
			title: {
				text: "PERIODE 2021",
			},
			subtitle: {
				text: "Source: TOTO",
			},
			xAxis: {
				categories: [
					"Jan",
					"Feb",
					"Mar",
					"Apr",
					"May",
					"Jun",
					"Jul",
					"Aug",
					"Sep",
					"Oct",
					"Nov",
					"Dec",
				],
				crosshair: true,
			},
			yAxis: {
				min: 0,
				title: {
					text: "Rainfall (mm)",
				},
			},
			tooltip: {
				headerFormat:
					'<span style="font-size:10px">{point.key}</span><table>',
				pointFormat:
					'<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f}</b></td></tr>',
				footerFormat: "</table>",
				shared: true,
				useHTML: true,
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0,
				},
			},
			series: pasaran,
		});
	}
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<Container fluid style="margin-top: 70px;">
	<Row>
		<Col xs="12">
			<div style="height:500px;" id="container" />
		</Col>
	</Row>
</Container>
