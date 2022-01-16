<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from "./components/Navigation.svelte";
	import Footer from "./components/Footer.svelte";
	import Home from "./pages/home/Home.svelte";
	import Pasaran from "./pages/pasaran/Pasaran.svelte";
	import Periode from "./pages/periode/Periode.svelte";
	import Prediksi from "./pages/prediksi/Prediksi.svelte";
	import Report_Winlose from "./pages/report/Winlose.svelte";
	import Admin from "./pages/admin/Admin.svelte";
	import AdminRule from "./pages/adminrule/Adminrule.svelte";
	import Log from "./pages/log/Log.svelte";
	import Login from "./pages/Login.svelte";
	import NotFound from "./pages/NotFound.svelte";
	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;

	if (token == "" || token == null) {
		routes = {
			"/": wrap({
				component: Login,
			}),
			"*": wrap({
				component: Login,
			}),
		};
	} else {
		isNav = true;
		routes = {
			"/": wrap({
				component: Home,
			}),
			"/log": wrap({
				component: Log,
			}),
			"/periode": wrap({
				component: Periode,
			}),
			"/pasaran": wrap({
				component: Pasaran,
			}),
			"/admin": wrap({
				component: Admin,
			}),
			"/adminrule": wrap({
				component: AdminRule,
			}),
			"/prediksi": wrap({
				component: Prediksi,
			}),
			"/reportwinlose": wrap({
				component: Report_Winlose,
			}),
			"*": NotFound,
		};
	}
</script>

{#if isNav}
	<Navigation />
{/if}
<Router {routes} />
<Footer />
