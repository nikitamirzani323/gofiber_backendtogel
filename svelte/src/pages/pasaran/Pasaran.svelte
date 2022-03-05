<script>
	import Home from "../pasaran/Home.svelte";
	import Entry from "../pasaran/Entry.svelte";

	let listPasaran = [];
	let listPasaranOnlineByEdit = [];
	let record = "";
	let totalrecord = 0;
	let idcomppasaran = "";
	let tipepasaran = "";
	let sData = "";
	let token = localStorage.getItem("token");
	let akses_page = true;
	let pasaran_idpasarantogel_field = "";
	let pasaran_name_field = "";
	let pasaran_tipepasaran_field = "";
	let pasaran_diundi_field = "";
	let pasaran_situs_field = "";
	let pasaran_tutup_field = "";
	let pasaran_jadwal_field = "";
	let pasaran_buka_field = "";
	let pasaran_limitline4d_field = 0;
	let pasaran_limitline3d_field = 0;
	let pasaran_limitline3dd_field = 0;
	let pasaran_limitline2d_field = 0;
	let pasaran_limitline2dd_field = 0;
	let pasaran_limitline2dt_field = 0;
	let pasaran_bbfs_field = 0;
	let pasaran_minbet_432d_field = 0;
	let pasaran_maxbet4d_432d_field = 0;
	let pasaran_maxbet3d_432d_field = 0;
	let pasaran_maxbet3dd_432d_field = 0;
	let pasaran_maxbet2d_432d_field = 0;
	let pasaran_maxbet2dd_432d_field = 0;
	let pasaran_maxbet2dt_432d_field = 0;
	let pasaran_limitotal4d_432d_field = 0;
	let pasaran_limitotal3d_432d_field = 0;
	let pasaran_limitotal3dd_432d_field = 0;
	let pasaran_limitotal2d_432d_field = 0;
	let pasaran_limitotal2dd_432d_field = 0;
	let pasaran_limitotal2dt_432d_field = 0;
	let pasaran_limitglobal4d_432d_field = 0;
	let pasaran_limitglobal3d_432d_field = 0;
	let pasaran_limitglobal3dd_432d_field = 0;
	let pasaran_limitglobal2d_432d_field = 0;
	let pasaran_limitglobal2dd_432d_field = 0;
	let pasaran_limitglobal2dt_432d_field = 0;
	let pasaran_disc4d_432d_field = 0;
	let pasaran_disc3d_432d_field = 0;
	let pasaran_disc3dd_432d_field = 0;
	let pasaran_disc2d_432d_field = 0;
	let pasaran_disc2dd_432d_field = 0;
	let pasaran_disc2dt_432d_field = 0;
	let pasaran_win4d_432d_field = 0;
	let pasaran_win3d_432d_field = 0;
	let pasaran_win3dd_432d_field = 0;
	let pasaran_win2d_432d_field = 0;
	let pasaran_win2dd_432d_field = 0;
	let pasaran_win2dt_432d_field = 0;
	let pasaran_win4dnodisc_432d_field = 0;
	let pasaran_win3dnodisc_432d_field = 0;
	let pasaran_win3ddnodisc_432d_field = 0;
	let pasaran_win2dnodisc_432d_field = 0;
	let pasaran_win2ddnodisc_432d_field = 0;
	let pasaran_win2dtnodisc_432d_field = 0;
	let pasaran_win4d_bb_kena_432d_field = 0;
    let pasaran_win3d_bb_kena_432d_field = 0;
    let pasaran_win3dd_bb_kena_432d_field = 0;
    let pasaran_win2d_bb_kena_432d_field = 0;
    let pasaran_win2dd_bb_kena_432d_field = 0;
    let pasaran_win2dt_bb_kena_432d_field = 0;
    let pasaran_win4d_bb_432d_field = 0;
    let pasaran_win3d_bb_432d_field = 0;
    let pasaran_win3dd_bb_432d_field = 0;
    let pasaran_win2d_bb_432d_field = 0;
    let pasaran_win2dd_bb_432d_field = 0;
    let pasaran_win2dt_bb_432d_field = 0;
	let pasaran_minbet_cbebas_field = 0;
	let pasaran_maxbet_cbebas_field = 0;
	let pasaran_limitotal_cbebas_field = 0;
	let pasaran_limitglobal_cbebas_field = 0;
	let pasaran_win_cbebas_field = 0;
	let pasaran_disc_cbebas_field = 0;
	let pasaran_minbet_cmacau_field = 0;
	let pasaran_maxbet_cmacau_field = 0;
	let pasaran_limitotal_cmacau_field = 0;
	let pasaran_limitglobal_cmacau_field = 0;
	let pasaran_win2_cmacau_field = 0;
	let pasaran_win3_cmacau_field = 0;
	let pasaran_win4_cmacau_field = 0;
	let pasaran_disc_cmacau_field = 0;
	let pasaran_minbet_cnaga_field = 0;
	let pasaran_maxbet_cnaga_field = 0;
	let pasaran_win3_cnaga_field = 0;
	let pasaran_win4_cnaga_field = 0;
	let pasaran_disc_cnaga_field = 0;
	let pasaran_limitglobal_cnaga_field = 0;
	let pasaran_limittotal_cnaga_field = 0;
	let pasaran_minbet_cjitu_field = 0;
	let pasaran_maxbet_cjitu_field = 0;
	let pasaran_winas_cjitu_field = 0;
	let pasaran_winkop_cjitu_field = 0;
	let pasaran_winkepala_cjitu_field = 0;
	let pasaran_winekor_cjitu_field = 0;
	let pasaran_desc_cjitu_field = 0;
	let pasaran_limitglobal_cjitu_field = 0;
	let pasaran_limittotal_cjitu_field = 0;
	let pasaran_minbet_5050umum_field = 0;
	let pasaran_maxbet_5050umum_field = 0;
	let pasaran_keibesar_5050umum_field = 0;
	let pasaran_keikecil_5050umum_field = 0;
	let pasaran_keigenap_5050umum_field = 0;
	let pasaran_keiganjil_5050umum_field = 0;
	let pasaran_keitengah_5050umum_field = 0;
	let pasaran_keitepi_5050umum_field = 0;
	let pasaran_discbesar_5050umum_field = 0;
	let pasaran_disckecil_5050umum_field = 0;
	let pasaran_discgenap_5050umum_field = 0;
	let pasaran_discganjil_5050umum_field = 0;
	let pasaran_disctengah_5050umum_field = 0;
	let pasaran_disctepi_5050umum_field = 0;
	let pasaran_limitglobal_5050umum_field = 0;
	let pasaran_limittotal_5050umum_field = 0;
	let pasaran_minbet_5050special_field = 0;
	let pasaran_maxbet_5050special_field = 0;
	let pasaran_keiasganjil_5050special_field = 0;
	let pasaran_keiasgenap_5050special_field = 0;
	let pasaran_keiasbesar_5050special_field = 0;
	let pasaran_keiaskecil_5050special_field = 0;
	let pasaran_keikopganjil_5050special_field = 0;
	let pasaran_keikopgenap_5050special_field = 0;
	let pasaran_keikopbesar_5050special_field = 0;
	let pasaran_keikopkecil_5050special_field = 0;
	let pasaran_keikepalaganjil_5050special_field = 0;
	let pasaran_keikepalagenap_5050special_field = 0;
	let pasaran_keikepalabesar_5050special_field = 0;
	let pasaran_keikepalakecil_5050special_field = 0;
	let pasaran_keiekorganjil_5050special_field = 0;
	let pasaran_keiekorgenap_5050special_field = 0;
	let pasaran_keiekorbesar_5050special_field = 0;
	let pasaran_keiekorkecil_5050special_field = 0;
	let pasaran_discasganjil_5050special_field = 0;
	let pasaran_discasgenap_5050special_field = 0;
	let pasaran_discasbesar_5050special_field = 0;
	let pasaran_discaskecil_5050special_field = 0;
	let pasaran_disckopganjil_5050special_field = 0;
	let pasaran_disckopgenap_5050special_field = 0;
	let pasaran_disckopbesar_5050special_field = 0;
	let pasaran_disckopkecil_5050special_field = 0;
	let pasaran_disckepalaganjil_5050special_field = 0;
	let pasaran_disckepalagenap_5050special_field = 0;
	let pasaran_disckepalabesar_5050special_field = 0;
	let pasaran_disckepalakecil_5050special_field = 0;
	let pasaran_discekorganjil_5050special_field = 0;
	let pasaran_discekorgenap_5050special_field = 0;
	let pasaran_discekorbesar_5050special_field = 0;
	let pasaran_discekorkecil_5050special_field = 0;
	let pasaran_limitglobal_5050special_field = 0;
	let pasaran_limittotal_5050special_field = 0;
	let pasaran_minbet_5050kombinasi_field = 0;
	let pasaran_maxbet_5050kombinasi_field = 0;
	let pasaran_belakangkeimono_5050kombinasi_field = 0;
	let pasaran_belakangkeistereo_5050kombinasi_field = 0;
	let pasaran_belakangkeikembang_5050kombinasi_field = 0;
	let pasaran_belakangkeikempis_5050kombinasi_field = 0;
	let pasaran_belakangkeikembar_5050kombinasi_field = 0;
	let pasaran_tengahkeimono_5050kombinasi_field = 0;
	let pasaran_tengahkeistereo_5050kombinasi_field = 0;
	let pasaran_tengahkeikembang_5050kombinasi_field = 0;
	let pasaran_tengahkeikempis_5050kombinasi_field = 0;
	let pasaran_tengahkeikembar_5050kombinasi_field = 0;
	let pasaran_depankeimono_5050kombinasi_field = 0;
	let pasaran_depankeistereo_5050kombinasi_field = 0;
	let pasaran_depankeikembang_5050kombinasi_field = 0;
	let pasaran_depankeikempis_5050kombinasi_field = 0;
	let pasaran_depankeikembar_5050kombinasi_field = 0;
	let pasaran_belakangdiscmono_5050kombinasi_field = 0;
	let pasaran_belakangdiscstereo_5050kombinasi_field = 0;
	let pasaran_belakangdisckembang_5050kombinasi_field = 0;
	let pasaran_belakangdisckempis_5050kombinasi_field = 0;
	let pasaran_belakangdisckembar_5050kombinasi_field = 0;
	let pasaran_tengahdiscmono_5050kombinasi_field = 0;
	let pasaran_tengahdiscstereo_5050kombinasi_field = 0;
	let pasaran_tengahdisckembang_5050kombinasi_field = 0;
	let pasaran_tengahdisckempis_5050kombinasi_field = 0;
	let pasaran_tengahdisckembar_5050kombinasi_field = 0;
	let pasaran_depandiscmono_5050kombinasi_field = 0;
	let pasaran_depandiscstereo_5050kombinasi_field = 0;
	let pasaran_depandisckembang_5050kombinasi_field = 0;
	let pasaran_depandisckempis_5050kombinasi_field = 0;
	let pasaran_depandisckembar_5050kombinasi_field = 0;
	let pasaran_limitglobal_5050kombinasi_field = 0;
	let pasaran_limittotal_5050kombinasi_field = 0;
	let pasaran_minbet_kombinasi_field = 0;
	let pasaran_maxbet_kombinasi_field = 0;
	let pasaran_win_kombinasi_field = 0;
	let pasaran_disc_kombinasi_field = 0;
	let pasaran_limitglobal_kombinasi_field = 0;
	let pasaran_limittotal_kombinasi_field = 0;
	let pasaran_minbet_dasar_field = 0;
	let pasaran_maxbet_dasar_field = 0;
	let pasaran_keibesar_dasar_field = 0;
	let pasaran_keikecil_dasar_field = 0;
	let pasaran_keigenap_dasar_field = 0;
	let pasaran_keiganjil_dasar_field = 0;
	let pasaran_discbesar_dasar_field = 0;
	let pasaran_disckecil_dasar_field = 0;
	let pasaran_discgenap_dasar_field = 0;
	let pasaran_discganjil_dasar_field = 0;
	let pasaran_limitglobal_dasar_field = 0;
	let pasaran_limittotal_dasar_field = 0;
	let pasaran_minbet_shio_field = 0;
	let pasaran_maxbet_shio_field = 0;
	let pasaran_win_shio_field = 0;
	let pasaran_disc_shio_field = 0;
	let pasaran_shioyear_shio_field = "";
	let pasaran_limitglobal_shio_field = 0;
	let pasaran_limittotal_shio_field = 0;
	let pasaran_display_field = "";
	let pasaran_status_field = "";
	let pasaran_create_field = "";
	let pasaran_createdate_field = "";
	let pasaran_update_field = "";
	let pasaran_updatedate_field = "";

	const handleRefreshData = (e) => {
		listPasaran = [];
		listPasaranOnlineByEdit = [];
		totalrecord = 0;
		setTimeout(function () {
			initPasaran();
		}, 1000);
	};
	const handleRefreshPasaranOnline = (e) => {
		let idcomppasaran = e.detail;
		listPasaranOnlineByEdit = [];
		editpasaran(idcomppasaran);
	};
	const handleEditData = (e) => {
		idcomppasaran = e.detail.e;
		tipepasaran = e.detail.f;
		sData = "Edit";
		listPasaranOnlineByEdit = []
		editpasaran(idcomppasaran,tipepasaran);
	};
	const handleDeletePasaranOnline = (e) => {
		let idcomppasaranoff = e.detail.e;
		let idcomppasaran = e.detail.idcomppasaran;
		deletePasaranOnline(idcomppasaranoff, idcomppasaran);
	};
	const handleBackHalaman = () => {
		idcomppasaran = "";
		sData = "";
		listPasaranOnlineByEdit = [];
		handleRefreshData("all");
	};
	async function deletePasaranOnline(e, z) {
		const res = await fetch("/api/deletepasaranonline", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				idpasaran: parseInt(z),
				idpasaraonline: parseInt(e),
				page:"PASARAN-SAVE",
			}),
		});
		const json = await res.json();

		if (json.status == 200) {
			listPasaranOnlineByEdit = [];
			editpasaran(z);
		}else if(json.status == 403){
			alert(json.message)
		}
	}
	async function editpasaran(e,f) {
		const res = await fetch("/api/editpasaran", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				idpasaran: e,
			}),
		});
		const json = await res.json();
		let record = json.record;
		let record_pasaranonline = json.pasaranonline;

		if (json.status === 400) {
			logout();
		} else {
			for (let i = 0; i < record.length; i++) {
				pasaran_idpasarantogel_field = record[i]["idpasarantogel"];
				pasaran_name_field = record[i]["nmpasaran"];
				pasaran_tipepasaran_field = f;
				pasaran_situs_field = record[i]["pasaranurl"];
				pasaran_diundi_field = record[i]["pasarandiundi"];
				pasaran_tutup_field = record[i]["jamtutup"];
				pasaran_jadwal_field = record[i]["jamjadwal"];
				pasaran_buka_field = record[i]["jamopen"];
				pasaran_limitline4d_field = record[i]["limitline_4d"];
				pasaran_limitline3d_field = record[i]["limitline_3d"];
				pasaran_limitline3dd_field = record[i]["limitline_3dd"];
				pasaran_limitline2d_field = record[i]["limitline_2d"];
				pasaran_limitline2dd_field = record[i]["limitline_2dd"];
				pasaran_limitline2dt_field = record[i]["limitline_2dt"];
				pasaran_bbfs_field = record[i]["bbfs"];
				pasaran_minbet_432d_field = record[i]["minbet_432d"];
				pasaran_maxbet4d_432d_field = record[i]["maxbet4d_432d"];
				pasaran_maxbet3d_432d_field = record[i]["maxbet3d_432d"];
				pasaran_maxbet3dd_432d_field = record[i]["maxbet3dd_432d"];
				pasaran_maxbet2d_432d_field = record[i]["maxbet2d_432d"];
				pasaran_maxbet2dd_432d_field = record[i]["maxbet2dd_432d"];
				pasaran_maxbet2dt_432d_field = record[i]["maxbet2dt_432d"];
				pasaran_limitotal4d_432d_field = record[i]["limitotal4d_432d"];
				pasaran_limitotal3d_432d_field = record[i]["limitotal3d_432d"];
				pasaran_limitotal3dd_432d_field = record[i]["limitotal3dd_432d"];
				pasaran_limitotal2d_432d_field = record[i]["limitotal2d_432d"];
				pasaran_limitotal2dd_432d_field = record[i]["limitotal2dd_432d"];
				pasaran_limitotal2dt_432d_field = record[i]["limitotal2dt_432d"];
				pasaran_limitglobal4d_432d_field = record[i]["limitglobal4d_432d"];
				pasaran_limitglobal3d_432d_field = record[i]["limitglobal3d_432d"];
				pasaran_limitglobal3dd_432d_field = record[i]["limitglobal3dd_432d"];
				pasaran_limitglobal2d_432d_field = record[i]["limitglobal2d_432d"];
				pasaran_limitglobal2dd_432d_field = record[i]["limitglobal2dd_432d"];
				pasaran_limitglobal2dt_432d_field = record[i]["limitglobal2dt_432d"];
				pasaran_disc4d_432d_field = (record[i]["disc4d_432d"] * 100).toFixed(2);
				pasaran_disc3d_432d_field = (record[i]["disc3d_432d"] * 100).toFixed(2);
				pasaran_disc3dd_432d_field = (record[i]["disc3dd_432d"] * 100).toFixed(2);
				pasaran_disc2d_432d_field = (record[i]["disc2d_432d"] * 100).toFixed(2);
				pasaran_disc2dd_432d_field = (record[i]["disc2dd_432d"] * 100).toFixed(2);
				pasaran_disc2dt_432d_field = (record[i]["disc2dt_432d"] * 100).toFixed(2);
				pasaran_win4d_432d_field = record[i]["win4d_432d"];
				pasaran_win3d_432d_field = record[i]["win3d_432d"];
				pasaran_win3dd_432d_field = record[i]["win3dd_432d"];
				pasaran_win2d_432d_field = record[i]["win2d_432d"];
				pasaran_win2dd_432d_field = record[i]["win2dd_432d"];
				pasaran_win2dt_432d_field = record[i]["win2dt_432d"];
				pasaran_win4dnodisc_432d_field = record[i]["win4dnodisc_432d"];
				pasaran_win3dnodisc_432d_field = record[i]["win3dnodisc_432d"];
				pasaran_win3ddnodisc_432d_field = record[i]["win3ddnodisc_432d"];
				pasaran_win2dnodisc_432d_field = record[i]["win2dnodisc_432d"];
				pasaran_win2ddnodisc_432d_field = record[i]["win2ddnodisc_432d"];
				pasaran_win2dtnodisc_432d_field = record[i]["win2dtnodisc_432d"];
				pasaran_win4d_bb_kena_432d_field = record[i]["win4dbb_kena_432d"];
                pasaran_win3d_bb_kena_432d_field = record[i]["win3dbb_kena_432d"];
                pasaran_win3dd_bb_kena_432d_field = record[i]["win3ddbb_kena_432d"];
                pasaran_win2d_bb_kena_432d_field = record[i]["win2dbb_kena_432d"];
                pasaran_win2dd_bb_kena_432d_field = record[i]["win2ddbb_kena_432d"];
                pasaran_win2dt_bb_kena_432d_field = record[i]["win2dtbb_kena_432d"];
                pasaran_win4d_bb_432d_field = record[i]["win4dbb_432d"];
                pasaran_win3d_bb_432d_field = record[i]["win3dbb_432d"];
                pasaran_win3dd_bb_432d_field = record[i]["win3ddbb_432d"];
                pasaran_win2d_bb_432d_field = record[i]["win2dbb_432d"];
                pasaran_win2dd_bb_432d_field = record[i]["win2ddbb_432d"];
                pasaran_win2dt_bb_432d_field = record[i]["win2dtbb_432d"];

				pasaran_minbet_cbebas_field = record[i]["minbet_cbebas"];
				pasaran_maxbet_cbebas_field = record[i]["maxbet_cbebas"];
				pasaran_limitotal_cbebas_field = record[i]["limittotal_cbebas"];
				pasaran_limitglobal_cbebas_field = record[i]["limitglobal_cbebas"];
				pasaran_win_cbebas_field = record[i]["win_cbebas"].toFixed(2);
				pasaran_disc_cbebas_field = (record[i]["disc_cbebas"] * 100).toFixed(2);

				pasaran_minbet_cmacau_field = record[i]["minbet_cmacau"];
				pasaran_maxbet_cmacau_field = record[i]["maxbet_cmacau"];
				pasaran_limitotal_cmacau_field = record[i]["limitotal_cmacau"];
				pasaran_limitglobal_cmacau_field = record[i]["limitglobal_cmacau"];
				pasaran_win2_cmacau_field = record[i]["win2d_cmacau"].toFixed(2);
				pasaran_win3_cmacau_field = record[i]["win3d_cmacau"].toFixed(2);
				pasaran_win4_cmacau_field = record[i]["win4d_cmacau"].toFixed(2);
				pasaran_disc_cmacau_field = (record[i]["disc_cmacau"] * 100).toFixed(2);

				pasaran_minbet_cnaga_field = record[i]["minbet_cnaga"];
				pasaran_maxbet_cnaga_field = record[i]["maxbet_cnaga"];
				pasaran_win3_cnaga_field = record[i]["win3_cnaga"].toFixed(2);
				pasaran_win4_cnaga_field = record[i]["win4_cnaga"].toFixed(2);
				pasaran_disc_cnaga_field = (record[i]["disc_cnaga"] * 100).toFixed(2);
				pasaran_limitglobal_cnaga_field = record[i]["limitglobal_cnaga"];
				pasaran_limittotal_cnaga_field = record[i]["limittotal_cnaga"];

				pasaran_minbet_cjitu_field = record[i]["minbet_cjitu"];
				pasaran_maxbet_cjitu_field = record[i]["maxbet_cjitu"];
				pasaran_winas_cjitu_field = record[i]["winas_cjitu"].toFixed(2);
				pasaran_winkop_cjitu_field = record[i]["winkop_cjitu"].toFixed(2);
				pasaran_winkepala_cjitu_field = record[i]["winkepala_cjitu"].toFixed(2);
				pasaran_winekor_cjitu_field = record[i]["winekor_cjitu"].toFixed(2);
				pasaran_desc_cjitu_field = (record[i]["desc_cjitu"] * 100).toFixed(2);
				pasaran_limitglobal_cjitu_field = record[i]["limitglobal_cjitu"];
				pasaran_limittotal_cjitu_field = record[i]["limittotal_cjitu"];

				pasaran_minbet_5050umum_field = record[i]["minbet_5050umum"];
				pasaran_maxbet_5050umum_field = record[i]["maxbet_5050umum"];
				pasaran_keibesar_5050umum_field = (record[i]["keibesar_5050umum"] * 100).toFixed(2);
				pasaran_keikecil_5050umum_field = (record[i]["keikecil_5050umum"] * 100).toFixed(2);
				pasaran_keigenap_5050umum_field = (record[i]["keigenap_5050umum"] * 100).toFixed(2);
				pasaran_keiganjil_5050umum_field = (record[i]["keiganjil_5050umum"] * 100).toFixed(2);
				pasaran_keitengah_5050umum_field = (record[i]["keitengah_5050umum"] * 100).toFixed(2);
				pasaran_keitepi_5050umum_field = (record[i]["keitepi_5050umum"] * 100).toFixed(2);
				pasaran_discbesar_5050umum_field = (record[i]["discbesar_5050umum"] * 100).toFixed(2);
				pasaran_disckecil_5050umum_field = (record[i]["disckecil_5050umum"] * 100).toFixed(2);
				pasaran_discgenap_5050umum_field = (record[i]["discgenap_5050umum"] * 100).toFixed(2);
				pasaran_discganjil_5050umum_field = (record[i]["discganjil_5050umum"] * 100).toFixed(2);
				pasaran_disctengah_5050umum_field = (record[i]["disctengah_5050umum"] * 100).toFixed(2);
				pasaran_disctepi_5050umum_field = (record[i]["disctepi_5050umum"] * 100).toFixed(2);
				pasaran_limitglobal_5050umum_field = record[i]["limitglobal_5050umum"];
				pasaran_limittotal_5050umum_field = record[i]["limittotal_5050umum"];

				pasaran_minbet_5050special_field = record[i]["minbet_5050special"];
				pasaran_maxbet_5050special_field = record[i]["maxbet_5050special"];
				pasaran_keiasganjil_5050special_field = (record[i]["keiasganjil_5050special"] * 100).toFixed(2);
				pasaran_keiasgenap_5050special_field = (record[i]["keiasgenap_5050special"] * 100).toFixed(2);
				pasaran_keiasbesar_5050special_field = (record[i]["keiasbesar_5050special"] * 100).toFixed(2);
				pasaran_keiaskecil_5050special_field = (record[i]["keiaskecil_5050special"] * 100).toFixed(2);
				pasaran_keikopganjil_5050special_field = (record[i]["keikopganjil_5050special"] * 100).toFixed(2);
				pasaran_keikopgenap_5050special_field = (record[i]["keikopgenap_5050special"] * 100).toFixed(2);
				pasaran_keikopbesar_5050special_field = (record[i]["keikopbesar_5050special"] * 100).toFixed(2);
				pasaran_keikopkecil_5050special_field = (record[i]["keikopkecil_5050special"] * 100).toFixed(2);
				pasaran_keikepalaganjil_5050special_field = (record[i]["keikepalaganjil_5050special"] * 100).toFixed(2);
				pasaran_keikepalagenap_5050special_field = (record[i]["keikepalagenap_5050special"] * 100).toFixed(2);
				pasaran_keikepalabesar_5050special_field = (record[i]["keikepalabesar_5050special"] * 100).toFixed(2);
				pasaran_keikepalakecil_5050special_field = (record[i]["keikepalakecil_5050special"] * 100).toFixed(2);
				pasaran_keiekorganjil_5050special_field = (record[i]["keiekorganjil_5050special"] * 100).toFixed(2);
				pasaran_keiekorgenap_5050special_field = (record[i]["keiekorgenap_5050special"] * 100).toFixed(2);
				pasaran_keiekorbesar_5050special_field = (record[i]["keiekorbesar_5050special"] * 100).toFixed(2);
				pasaran_keiekorkecil_5050special_field = (record[i]["keiekorkecil_5050special"] * 100).toFixed(2);
				pasaran_discasganjil_5050special_field = (record[i]["discasganjil_5050special"] * 100).toFixed(2);
				pasaran_discasgenap_5050special_field = (record[i]["discasgenap_5050special"] * 100).toFixed(2);
				pasaran_discasbesar_5050special_field = (record[i]["discasbesar_5050special"] * 100).toFixed(2);
				pasaran_discaskecil_5050special_field = (record[i]["discaskecil_5050special"] * 100).toFixed(2);
				pasaran_disckopganjil_5050special_field = (record[i]["disckopganjil_5050special"] * 100).toFixed(2);
				pasaran_disckopgenap_5050special_field = (record[i]["disckopgenap_5050special"] * 100).toFixed(2);
				pasaran_disckopbesar_5050special_field = (record[i]["disckopbesar_5050special"] * 100).toFixed(2);
				pasaran_disckopkecil_5050special_field = (record[i]["disckopkecil_5050special"] * 100).toFixed(2);
				pasaran_disckepalaganjil_5050special_field = (record[i]["disckepalaganjil_5050special"] * 100).toFixed(2);
				pasaran_disckepalagenap_5050special_field = (record[i]["disckepalagenap_5050special"] * 100).toFixed(2);
				pasaran_disckepalabesar_5050special_field = (record[i]["disckepalabesar_5050special"] * 100).toFixed(2);
				pasaran_disckepalakecil_5050special_field = (record[i]["disckepalakecil_5050special"] * 100).toFixed(2);
				pasaran_discekorganjil_5050special_field = (record[i]["discekorganjil_5050special"] * 100).toFixed(2);
				pasaran_discekorgenap_5050special_field = (record[i]["discekorgenap_5050special"] * 100).toFixed(2);
				pasaran_discekorbesar_5050special_field = (record[i]["discekorbesar_5050special"] * 100).toFixed(2);
				pasaran_discekorkecil_5050special_field = (record[i]["discekorkecil_5050special"] * 100).toFixed(2);
				pasaran_limitglobal_5050special_field = record[i]["limitglobal_5050special"];
				pasaran_limittotal_5050special_field = record[i]["limittotal_5050special"];

				pasaran_minbet_5050kombinasi_field = record[i]["minbet_5050kombinasi"];
				pasaran_maxbet_5050kombinasi_field = record[i]["maxbet_5050kombinasi"];
				pasaran_belakangkeimono_5050kombinasi_field = (record[i]["belakangkeimono_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangkeistereo_5050kombinasi_field = (record[i]["belakangkeistereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangkeikembang_5050kombinasi_field = (record[i]["belakangkeikembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangkeikempis_5050kombinasi_field = (record[i]["belakangkeikempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangkeikembar_5050kombinasi_field = (record[i]["belakangkeikembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahkeimono_5050kombinasi_field = (record[i]["tengahkeimono_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahkeistereo_5050kombinasi_field = (record[i]["tengahkeistereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahkeikembang_5050kombinasi_field = (record[i]["tengahkeikembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahkeikempis_5050kombinasi_field = (record[i]["tengahkeikempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahkeikembar_5050kombinasi_field = (record[i]["tengahkeikembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_depankeimono_5050kombinasi_field = (record[i]["depankeimono_5050kombinasi"] * 100).toFixed(2);
				pasaran_depankeistereo_5050kombinasi_field = (record[i]["depankeistereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_depankeikembang_5050kombinasi_field = (record[i]["depankeikembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_depankeikempis_5050kombinasi_field = (record[i]["depankeikempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_depankeikembar_5050kombinasi_field = (record[i]["depankeikembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangdiscmono_5050kombinasi_field = (record[i]["belakangdiscmono_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangdiscstereo_5050kombinasi_field = (record[i]["belakangdiscstereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangdisckembang_5050kombinasi_field = (record[i]["belakangdisckembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangdisckempis_5050kombinasi_field = (record[i]["belakangdisckempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_belakangdisckembar_5050kombinasi_field = (record[i]["belakangdisckembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahdiscmono_5050kombinasi_field = (record[i]["tengahdiscmono_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahdiscstereo_5050kombinasi_field = (record[i]["tengahdiscstereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahdisckembang_5050kombinasi_field = (record[i]["tengahdisckembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahdisckempis_5050kombinasi_field = (record[i]["tengahdisckempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_tengahdisckembar_5050kombinasi_field = (record[i]["tengahdisckembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_depandiscmono_5050kombinasi_field = (record[i]["depandiscmono_5050kombinasi"] * 100).toFixed(2);
				pasaran_depandiscstereo_5050kombinasi_field = (record[i]["depandiscstereo_5050kombinasi"] * 100).toFixed(2);
				pasaran_depandisckembang_5050kombinasi_field = (record[i]["depandisckembang_5050kombinasi"] * 100).toFixed(2);
				pasaran_depandisckempis_5050kombinasi_field = (record[i]["depandisckempis_5050kombinasi"] * 100).toFixed(2);
				pasaran_depandisckembar_5050kombinasi_field = (record[i]["depandisckembar_5050kombinasi"] * 100).toFixed(2);
				pasaran_limitglobal_5050kombinasi_field = record[i]["limitglobal_5050kombinasi"];
				pasaran_limittotal_5050kombinasi_field = record[i]["limittotal_5050kombinasi"];

				pasaran_minbet_kombinasi_field = record[i]["minbet_kombinasi"];
				pasaran_maxbet_kombinasi_field = record[i]["maxbet_kombinasi"];
				pasaran_win_kombinasi_field = record[i]["win_kombinasi"].toFixed(2);
				pasaran_disc_kombinasi_field = (record[i]["disc_kombinasi"] * 100).toFixed(2);
				pasaran_limitglobal_kombinasi_field = record[i]["limitglobal_kombinasi"];
				pasaran_limittotal_kombinasi_field = record[i]["limittotal_kombinasi"];

				pasaran_minbet_dasar_field = record[i]["minbet_dasar"];
				pasaran_maxbet_dasar_field = record[i]["maxbet_dasar"];
				pasaran_keibesar_dasar_field = (record[i]["keibesar_dasar"] * 100).toFixed(2);
				pasaran_keikecil_dasar_field = (record[i]["keikecil_dasar"] * 100).toFixed(2);
				pasaran_keigenap_dasar_field = (record[i]["keigenap_dasar"] * 100).toFixed(2);
				pasaran_keiganjil_dasar_field = (record[i]["keiganjil_dasar"] * 100).toFixed(2);
				pasaran_discbesar_dasar_field = (record[i]["discbesar_dasar"] * 100).toFixed(2);
				pasaran_disckecil_dasar_field = (record[i]["disckecil_dasar"] * 100).toFixed(2);
				pasaran_discgenap_dasar_field = (record[i]["discgenap_dasar"] * 100).toFixed(2);
				pasaran_discganjil_dasar_field = (record[i]["discganjil_dasar"] * 100).toFixed(2);
				pasaran_limitglobal_dasar_field = record[i]["limitglobal_dasar"];
				pasaran_limittotal_dasar_field = record[i]["limittotal_dasar"];
				
				pasaran_minbet_shio_field = record[i]["minbet_shio"];
				pasaran_maxbet_shio_field = record[i]["maxbet_shio"];
				pasaran_win_shio_field = record[i]["win_shio"].toFixed(2);
				pasaran_disc_shio_field = (record[i]["disc_shio"] * 100).toFixed(2);
				pasaran_shioyear_shio_field = record[i]["shioyear_shio"];
				pasaran_limitglobal_shio_field = record[i]["limitglobal_shio"];
				pasaran_limittotal_shio_field = record[i]["limittotal_shio"];

				pasaran_display_field = record[i]["displaypasaran"];
				pasaran_status_field = record[i]["statuspasaranactive"];
				pasaran_create_field = record[i]["create"];
				pasaran_createdate_field = record[i]["createdate"];
				pasaran_update_field = record[i]["update"];
				pasaran_updatedate_field = record[i]["updatedate"];
			}
			for (var i = 0; i < record_pasaranonline.length; i++) {
				listPasaranOnlineByEdit = [
					...listPasaranOnlineByEdit,
					{
						idpasaranonline: record_pasaranonline[i]["Idpasaranonline"],
						haripasaran: capitalizeFirstLetter(record_pasaranonline[i]["haripasaran"]),
					},
				];
			}
		}
	}
	function capitalizeFirstLetter(string) {
		return string.charAt(0).toUpperCase() + string.slice(1);
	}
	async function initapp() {
		const res = await fetch("/api/home", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				page: "PASARAN-VIEW",
			}),
		});
		const json = await res.json();
		if (json.status === 400) {
			logout();
		} else if (json.status == 403) {
			alert(json.message);
			akses_page = false;
		} else {
			initPasaran();
		}
	}

	async function initPasaran() {
		const res = await fetch("/api/allpasaran", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
			}),
		});
		const json = await res.json();
		if (json.status == 200) {
			record = json.record;
			totalrecord = record.length;
			if (record != null) {
				for (var i = 0; i < record.length; i++) {
					listPasaran = [
						...listPasaran,
						{
							idcomppasaran: record[i]["idcomppasaran"],
							nmpasarantogel: record[i]["nmpasarantogel"],
							tipepasaran: record[i]["tipepasaran"],
							pasarandiundi: record[i]["pasarandiundi"],
							jamtutup: record[i]["jamtutup"],
							jamjadwal: record[i]["jamjadwal"],
							jamopen: record[i]["jamopen"],
							displaypasaran: record[i]["displaypasaran"],
							statuspasaran: record[i]["statuspasaran"],
							statuspasaranactive: record[i]["statuspasaranactive"],
							statuspasaran_css: record[i]["statuspasaran_css"],
							statuspasaranactive_css: record[i]["statuspasaranactive_css"],
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
	async function logout() {
		localStorage.clear();
		window.location.href = "/";
	}
	initapp();
</script>

{#if sData == "" && idcomppasaran == ""}
	{#if akses_page == true}
		<Home
			on:handleRefreshData={handleRefreshData}
			on:handleEditData={handleEditData}
			{listPasaran}
			{totalrecord}
		/>
	{/if}
{:else}
	<Entry
		on:handleBackHalaman={handleBackHalaman}
		on:handleRefreshPasaranOnline={handleRefreshPasaranOnline}
		on:handleDeletePasaranOnline={handleDeletePasaranOnline}
		{sData}
		{idcomppasaran}
		{token}
		{pasaran_idpasarantogel_field}
		{pasaran_name_field}
		{pasaran_tipepasaran_field}
		{pasaran_situs_field}
		{pasaran_diundi_field}
		{pasaran_tutup_field}
		{pasaran_jadwal_field}
		{pasaran_buka_field}
		{pasaran_limitline4d_field}
		{pasaran_limitline3d_field}
		{pasaran_limitline3dd_field}
		{pasaran_limitline2d_field}
		{pasaran_limitline2dd_field}
		{pasaran_limitline2dt_field}
		{pasaran_bbfs_field}
		{pasaran_minbet_432d_field}
		{pasaran_maxbet4d_432d_field}
		{pasaran_maxbet3d_432d_field}
		{pasaran_maxbet3dd_432d_field}
		{pasaran_maxbet2d_432d_field}
		{pasaran_maxbet2dd_432d_field}
		{pasaran_maxbet2dt_432d_field}
		{pasaran_limitotal4d_432d_field}
		{pasaran_limitotal3d_432d_field}
		{pasaran_limitotal3dd_432d_field}
		{pasaran_limitotal2d_432d_field}
		{pasaran_limitotal2dd_432d_field}
		{pasaran_limitotal2dt_432d_field}
		{pasaran_limitglobal4d_432d_field}
		{pasaran_limitglobal3d_432d_field}
		{pasaran_limitglobal3dd_432d_field}
		{pasaran_limitglobal2d_432d_field}
		{pasaran_limitglobal2dd_432d_field}
		{pasaran_limitglobal2dt_432d_field}
		{pasaran_disc4d_432d_field}
		{pasaran_disc3d_432d_field}
		{pasaran_disc3dd_432d_field}
		{pasaran_disc2d_432d_field}
		{pasaran_disc2dd_432d_field}
		{pasaran_disc2dt_432d_field}
		{pasaran_win4d_432d_field}
		{pasaran_win3d_432d_field}
		{pasaran_win3dd_432d_field}
		{pasaran_win2d_432d_field}
		{pasaran_win2dd_432d_field}
		{pasaran_win2dt_432d_field}
		{pasaran_win4dnodisc_432d_field}
		{pasaran_win3dnodisc_432d_field}
		{pasaran_win3ddnodisc_432d_field}
		{pasaran_win2dnodisc_432d_field}
		{pasaran_win2ddnodisc_432d_field}
		{pasaran_win2dtnodisc_432d_field}
		{pasaran_win4d_bb_kena_432d_field}
		{pasaran_win3d_bb_kena_432d_field}
		{pasaran_win3dd_bb_kena_432d_field}
		{pasaran_win2d_bb_kena_432d_field}
		{pasaran_win2dd_bb_kena_432d_field}
		{pasaran_win2dt_bb_kena_432d_field}
		{pasaran_win4d_bb_432d_field}
		{pasaran_win3d_bb_432d_field}
		{pasaran_win3dd_bb_432d_field}
		{pasaran_win2d_bb_432d_field}
		{pasaran_win2dd_bb_432d_field}
		{pasaran_win2dt_bb_432d_field}
		{pasaran_minbet_cbebas_field}
		{pasaran_maxbet_cbebas_field}
		{pasaran_limitotal_cbebas_field}
		{pasaran_limitglobal_cbebas_field}
		{pasaran_win_cbebas_field}
		{pasaran_disc_cbebas_field}
		{pasaran_minbet_cmacau_field}
		{pasaran_maxbet_cmacau_field}
		{pasaran_limitotal_cmacau_field}
		{pasaran_limitglobal_cmacau_field}
		{pasaran_win2_cmacau_field}
		{pasaran_win3_cmacau_field}
		{pasaran_win4_cmacau_field}
		{pasaran_disc_cmacau_field}
		{pasaran_minbet_cnaga_field}
		{pasaran_maxbet_cnaga_field}
		{pasaran_win3_cnaga_field}
		{pasaran_win4_cnaga_field}
		{pasaran_disc_cnaga_field}
		{pasaran_limitglobal_cnaga_field}
		{pasaran_limittotal_cnaga_field}
		{pasaran_minbet_cjitu_field}
		{pasaran_maxbet_cjitu_field}
		{pasaran_winas_cjitu_field}
		{pasaran_winkop_cjitu_field}
		{pasaran_winkepala_cjitu_field}
		{pasaran_winekor_cjitu_field}
		{pasaran_desc_cjitu_field}
		{pasaran_limitglobal_cjitu_field}
		{pasaran_limittotal_cjitu_field}
		{pasaran_minbet_5050umum_field}
		{pasaran_maxbet_5050umum_field}
		{pasaran_keibesar_5050umum_field}
		{pasaran_keikecil_5050umum_field}
		{pasaran_keigenap_5050umum_field}
		{pasaran_keiganjil_5050umum_field}
		{pasaran_keitengah_5050umum_field}
		{pasaran_keitepi_5050umum_field}
		{pasaran_discbesar_5050umum_field}
		{pasaran_disckecil_5050umum_field}
		{pasaran_discgenap_5050umum_field}
		{pasaran_discganjil_5050umum_field}
		{pasaran_disctengah_5050umum_field}
		{pasaran_disctepi_5050umum_field}
		{pasaran_limitglobal_5050umum_field}
		{pasaran_limittotal_5050umum_field}
		{pasaran_minbet_5050special_field}
		{pasaran_maxbet_5050special_field}
		{pasaran_keiasganjil_5050special_field}
		{pasaran_keiasgenap_5050special_field}
		{pasaran_keiasbesar_5050special_field}
		{pasaran_keiaskecil_5050special_field}
		{pasaran_keikopganjil_5050special_field}
		{pasaran_keikopgenap_5050special_field}
		{pasaran_keikopbesar_5050special_field}
		{pasaran_keikopkecil_5050special_field}
		{pasaran_keikepalaganjil_5050special_field}
		{pasaran_keikepalagenap_5050special_field}
		{pasaran_keikepalabesar_5050special_field}
		{pasaran_keikepalakecil_5050special_field}
		{pasaran_keiekorganjil_5050special_field}
		{pasaran_keiekorgenap_5050special_field}
		{pasaran_keiekorbesar_5050special_field}
		{pasaran_keiekorkecil_5050special_field}
		{pasaran_discasganjil_5050special_field}
		{pasaran_discasgenap_5050special_field}
		{pasaran_discasbesar_5050special_field}
		{pasaran_discaskecil_5050special_field}
		{pasaran_disckopganjil_5050special_field}
		{pasaran_disckopgenap_5050special_field}
		{pasaran_disckopbesar_5050special_field}
		{pasaran_disckopkecil_5050special_field}
		{pasaran_disckepalaganjil_5050special_field}
		{pasaran_disckepalagenap_5050special_field}
		{pasaran_disckepalabesar_5050special_field}
		{pasaran_disckepalakecil_5050special_field}
		{pasaran_discekorganjil_5050special_field}
		{pasaran_discekorgenap_5050special_field}
		{pasaran_discekorbesar_5050special_field}
		{pasaran_discekorkecil_5050special_field}
		{pasaran_limitglobal_5050special_field}
		{pasaran_limittotal_5050special_field}
		{pasaran_minbet_5050kombinasi_field}
		{pasaran_maxbet_5050kombinasi_field}
		{pasaran_belakangkeimono_5050kombinasi_field}
		{pasaran_belakangkeistereo_5050kombinasi_field}
		{pasaran_belakangkeikembang_5050kombinasi_field}
		{pasaran_belakangkeikempis_5050kombinasi_field}
		{pasaran_belakangkeikembar_5050kombinasi_field}
		{pasaran_tengahkeimono_5050kombinasi_field}
		{pasaran_tengahkeistereo_5050kombinasi_field}
		{pasaran_tengahkeikembang_5050kombinasi_field}
		{pasaran_tengahkeikempis_5050kombinasi_field}
		{pasaran_tengahkeikembar_5050kombinasi_field}
		{pasaran_depankeimono_5050kombinasi_field}
		{pasaran_depankeistereo_5050kombinasi_field}
		{pasaran_depankeikembang_5050kombinasi_field}
		{pasaran_depankeikempis_5050kombinasi_field}
		{pasaran_depankeikembar_5050kombinasi_field}
		{pasaran_belakangdiscmono_5050kombinasi_field}
		{pasaran_belakangdiscstereo_5050kombinasi_field}
		{pasaran_belakangdisckembang_5050kombinasi_field}
		{pasaran_belakangdisckempis_5050kombinasi_field}
		{pasaran_belakangdisckembar_5050kombinasi_field}
		{pasaran_tengahdiscmono_5050kombinasi_field}
		{pasaran_tengahdiscstereo_5050kombinasi_field}
		{pasaran_tengahdisckembang_5050kombinasi_field}
		{pasaran_tengahdisckempis_5050kombinasi_field}
		{pasaran_tengahdisckembar_5050kombinasi_field}
		{pasaran_depandiscmono_5050kombinasi_field}
		{pasaran_depandiscstereo_5050kombinasi_field}
		{pasaran_depandisckembang_5050kombinasi_field}
		{pasaran_depandisckempis_5050kombinasi_field}
		{pasaran_depandisckembar_5050kombinasi_field}
		{pasaran_limitglobal_5050kombinasi_field}
		{pasaran_limittotal_5050kombinasi_field}
		{pasaran_minbet_kombinasi_field}
		{pasaran_maxbet_kombinasi_field}
		{pasaran_win_kombinasi_field}
		{pasaran_disc_kombinasi_field}
		{pasaran_limitglobal_kombinasi_field}
		{pasaran_limittotal_kombinasi_field}
		{pasaran_minbet_dasar_field}
		{pasaran_maxbet_dasar_field}
		{pasaran_keibesar_dasar_field}
		{pasaran_keikecil_dasar_field}
		{pasaran_keigenap_dasar_field}
		{pasaran_keiganjil_dasar_field}
		{pasaran_discbesar_dasar_field}
		{pasaran_disckecil_dasar_field}
		{pasaran_discgenap_dasar_field}
		{pasaran_discganjil_dasar_field}
		{pasaran_limitglobal_dasar_field}
		{pasaran_limittotal_dasar_field}
		{pasaran_minbet_shio_field}
		{pasaran_maxbet_shio_field}
		{pasaran_win_shio_field}
		{pasaran_disc_shio_field}
		{pasaran_shioyear_shio_field}
		{pasaran_limitglobal_shio_field}
		{pasaran_limittotal_shio_field}
		{pasaran_display_field}
		{pasaran_status_field}
		{pasaran_create_field}
		{pasaran_createdate_field}
		{pasaran_update_field}
		{pasaran_updatedate_field}
		{listPasaranOnlineByEdit}
	/>
{/if}
