<script>
    import {
        Icon,
        Row,
        Col,
        Container,
        TabContent,
        TabPane,
    } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    import { createEventDispatcher } from "svelte";

    export let sData = "";
    export let idcomppasaran = "";
    export let token = "";
    export let pasaran_idpasarantogel_field = "";
    export let pasaran_name_field = "";
    export let pasaran_tipepasaran_field = "";
    export let pasaran_situs_field = "";
    export let pasaran_diundi_field = "";
    export let pasaran_tutup_field = "";
    export let pasaran_jadwal_field = "";
    export let pasaran_buka_field = "";
    export let pasaran_limitline4d_field = 0;
    export let pasaran_limitline3d_field = 0;
    export let pasaran_limitline3dd_field = 0;
    export let pasaran_limitline2d_field = 0;
    export let pasaran_limitline2dd_field = 0;
    export let pasaran_limitline2dt_field = 0;
    export let pasaran_bbfs_field = 0;
    export let pasaran_minbet_432d_field = 0;
    export let pasaran_maxbet4d_432d_field = 0;
    export let pasaran_maxbet3d_432d_field = 0;
    export let pasaran_maxbet3dd_432d_field = 0;
    export let pasaran_maxbet2d_432d_field = 0;
    export let pasaran_maxbet2dd_432d_field = 0;
    export let pasaran_maxbet2dt_432d_field = 0;
    export let pasaran_limitotal4d_432d_field = 0;
    export let pasaran_limitotal3d_432d_field = 0;
    export let pasaran_limitotal3dd_432d_field = 0;
    export let pasaran_limitotal2d_432d_field = 0;
    export let pasaran_limitotal2dd_432d_field = 0;
    export let pasaran_limitotal2dt_432d_field = 0;
    export let pasaran_limitglobal4d_432d_field = 0;
    export let pasaran_limitglobal3d_432d_field = 0;
    export let pasaran_limitglobal3dd_432d_field = 0;
    export let pasaran_limitglobal2d_432d_field = 0;
    export let pasaran_limitglobal2dd_432d_field = 0;
    export let pasaran_limitglobal2dt_432d_field = 0;
    export let pasaran_disc4d_432d_field = 0;
    export let pasaran_disc3d_432d_field = 0;
    export let pasaran_disc3dd_432d_field = 0;
    export let pasaran_disc2d_432d_field = 0;
    export let pasaran_disc2dd_432d_field = 0;
    export let pasaran_disc2dt_432d_field = 0;
    export let pasaran_win4d_432d_field = 0;
    export let pasaran_win3d_432d_field = 0;
    export let pasaran_win3dd_432d_field = 0;
    export let pasaran_win2d_432d_field = 0;
    export let pasaran_win2dd_432d_field = 0;
    export let pasaran_win2dt_432d_field = 0;

    export let pasaran_minbet_cbebas_field = 0;
    export let pasaran_maxbet_cbebas_field = 0;
    export let pasaran_limitotal_cbebas_field = 0;
    export let pasaran_limitglobal_cbebas_field = 0;
    export let pasaran_win_cbebas_field = 0;
    export let pasaran_disc_cbebas_field = 0;

    export let pasaran_minbet_cmacau_field = 0;
    export let pasaran_maxbet_cmacau_field = 0;
    export let pasaran_limitotal_cmacau_field = 0;
    export let pasaran_limitglobal_cmacau_field = 0;
    export let pasaran_win2_cmacau_field = 0;
    export let pasaran_win3_cmacau_field = 0;
    export let pasaran_win4_cmacau_field = 0;
    export let pasaran_disc_cmacau_field = 0;

    export let pasaran_minbet_cnaga_field = 0;
    export let pasaran_maxbet_cnaga_field = 0;
    export let pasaran_win3_cnaga_field = 0;
    export let pasaran_win4_cnaga_field = 0;
    export let pasaran_disc_cnaga_field = 0;
    export let pasaran_limitglobal_cnaga_field = 0;
    export let pasaran_limittotal_cnaga_field = 0;

    export let pasaran_minbet_cjitu_field = 0;
    export let pasaran_maxbet_cjitu_field = 0;
    export let pasaran_winas_cjitu_field = 0;
    export let pasaran_winkop_cjitu_field = 0;
    export let pasaran_winkepala_cjitu_field = 0;
    export let pasaran_winekor_cjitu_field = 0;
    export let pasaran_desc_cjitu_field = 0;
    export let pasaran_limitglobal_cjitu_field = 0;
    export let pasaran_limittotal_cjitu_field = 0;

    export let pasaran_minbet_5050umum_field = 0;
    export let pasaran_maxbet_5050umum_field = 0;
    export let pasaran_keibesar_5050umum_field = 0;
    export let pasaran_keikecil_5050umum_field = 0;
    export let pasaran_keigenap_5050umum_field = 0;
    export let pasaran_keiganjil_5050umum_field = 0;
    export let pasaran_keitengah_5050umum_field = 0;
    export let pasaran_keitepi_5050umum_field = 0;
    export let pasaran_discbesar_5050umum_field = 0;
    export let pasaran_disckecil_5050umum_field = 0;
    export let pasaran_discgenap_5050umum_field = 0;
    export let pasaran_discganjil_5050umum_field = 0;
    export let pasaran_disctengah_5050umum_field = 0;
    export let pasaran_disctepi_5050umum_field = 0;
    export let pasaran_limitglobal_5050umum_field = 0;
    export let pasaran_limittotal_5050umum_field = 0;

    export let pasaran_minbet_5050special_field = 0;
    export let pasaran_maxbet_5050special_field = 0;
    export let pasaran_keiasganjil_5050special_field = 0;
    export let pasaran_keiasgenap_5050special_field = 0;
    export let pasaran_keiasbesar_5050special_field = 0;
    export let pasaran_keiaskecil_5050special_field = 0;
    export let pasaran_keikopganjil_5050special_field = 0;
    export let pasaran_keikopgenap_5050special_field = 0;
    export let pasaran_keikopbesar_5050special_field = 0;
    export let pasaran_keikopkecil_5050special_field = 0;
    export let pasaran_keikepalaganjil_5050special_field = 0;
    export let pasaran_keikepalagenap_5050special_field = 0;
    export let pasaran_keikepalabesar_5050special_field = 0;
    export let pasaran_keikepalakecil_5050special_field = 0;
    export let pasaran_keiekorganjil_5050special_field = 0;
    export let pasaran_keiekorgenap_5050special_field = 0;
    export let pasaran_keiekorbesar_5050special_field = 0;
    export let pasaran_keiekorkecil_5050special_field = 0;
    export let pasaran_discasganjil_5050special_field = 0;
    export let pasaran_discasgenap_5050special_field = 0;
    export let pasaran_discasbesar_5050special_field = 0;
    export let pasaran_discaskecil_5050special_field = 0;
    export let pasaran_disckopganjil_5050special_field = 0;
    export let pasaran_disckopgenap_5050special_field = 0;
    export let pasaran_disckopbesar_5050special_field = 0;
    export let pasaran_disckopkecil_5050special_field = 0;
    export let pasaran_disckepalaganjil_5050special_field = 0;
    export let pasaran_disckepalagenap_5050special_field = 0;
    export let pasaran_disckepalabesar_5050special_field = 0;
    export let pasaran_disckepalakecil_5050special_field = 0;
    export let pasaran_discekorganjil_5050special_field = 0;
    export let pasaran_discekorgenap_5050special_field = 0;
    export let pasaran_discekorbesar_5050special_field = 0;
    export let pasaran_discekorkecil_5050special_field = 0;
    export let pasaran_limitglobal_5050special_field = 0;
    export let pasaran_limittotal_5050special_field = 0;

    export let pasaran_minbet_5050kombinasi_field = 0;
    export let pasaran_maxbet_5050kombinasi_field = 0;
    export let pasaran_belakangkeimono_5050kombinasi_field = 0;
    export let pasaran_belakangkeistereo_5050kombinasi_field = 0;
    export let pasaran_belakangkeikembang_5050kombinasi_field = 0;
    export let pasaran_belakangkeikempis_5050kombinasi_field = 0;
    export let pasaran_belakangkeikembar_5050kombinasi_field = 0;
    export let pasaran_tengahkeimono_5050kombinasi_field = 0;
    export let pasaran_tengahkeistereo_5050kombinasi_field = 0;
    export let pasaran_tengahkeikembang_5050kombinasi_field = 0;
    export let pasaran_tengahkeikempis_5050kombinasi_field = 0;
    export let pasaran_tengahkeikembar_5050kombinasi_field = 0;
    export let pasaran_depankeimono_5050kombinasi_field = 0;
    export let pasaran_depankeistereo_5050kombinasi_field = 0;
    export let pasaran_depankeikembang_5050kombinasi_field = 0;
    export let pasaran_depankeikempis_5050kombinasi_field = 0;
    export let pasaran_depankeikembar_5050kombinasi_field = 0;
    export let pasaran_belakangdiscmono_5050kombinasi_field = 0;
    export let pasaran_belakangdiscstereo_5050kombinasi_field = 0;
    export let pasaran_belakangdisckembang_5050kombinasi_field = 0;
    export let pasaran_belakangdisckempis_5050kombinasi_field = 0;
    export let pasaran_belakangdisckembar_5050kombinasi_field = 0;
    export let pasaran_tengahdiscmono_5050kombinasi_field = 0;
    export let pasaran_tengahdiscstereo_5050kombinasi_field = 0;
    export let pasaran_tengahdisckembang_5050kombinasi_field = 0;
    export let pasaran_tengahdisckempis_5050kombinasi_field = 0;
    export let pasaran_tengahdisckembar_5050kombinasi_field = 0;
    export let pasaran_depandiscmono_5050kombinasi_field = 0;
    export let pasaran_depandiscstereo_5050kombinasi_field = 0;
    export let pasaran_depandisckembang_5050kombinasi_field = 0;
    export let pasaran_depandisckempis_5050kombinasi_field = 0;
    export let pasaran_depandisckembar_5050kombinasi_field = 0;
    export let pasaran_limitglobal_5050kombinasi_field = 0;
    export let pasaran_limittotal_5050kombinasi_field = 0;

    export let pasaran_minbet_kombinasi_field = 0;
    export let pasaran_maxbet_kombinasi_field = 0;
    export let pasaran_win_kombinasi_field = 0;
    export let pasaran_disc_kombinasi_field = 0;
    export let pasaran_limitglobal_kombinasi_field = 0;
    export let pasaran_limittotal_kombinasi_field = 0;
    export let pasaran_minbet_dasar_field = 0;
    export let pasaran_maxbet_dasar_field = 0;
    export let pasaran_keibesar_dasar_field = 0;
    export let pasaran_keikecil_dasar_field = 0;
    export let pasaran_keigenap_dasar_field = 0;
    export let pasaran_keiganjil_dasar_field = 0;
    export let pasaran_discbesar_dasar_field = 0;
    export let pasaran_disckecil_dasar_field = 0;
    export let pasaran_discgenap_dasar_field = 0;
    export let pasaran_discganjil_dasar_field = 0;
    export let pasaran_limitglobal_dasar_field = 0;
    export let pasaran_limittotal_dasar_field = 0;
    export let pasaran_minbet_shio_field = 0;
    export let pasaran_maxbet_shio_field = 0;
    export let pasaran_win_shio_field = 0;
    export let pasaran_disc_shio_field = 0;
    export let pasaran_shioyear_shio_field = "";
    export let pasaran_limitglobal_shio_field = 0;
    export let pasaran_limittotal_shio_field = 0;

    export let pasaran_display_field = 0;
    export let pasaran_status_field = "";
    export let pasaran_create_field = "";
    export let pasaran_createdate_field = "";
    export let pasaran_update_field = "";
    export let pasaran_updatedate_field = "";
    export let listPasaranOnlineByEdit = [];
    let css_loader = "display: none;";
    let msgloader = "";
    let select_pasaranonline = "";
    let dispatch = createEventDispatcher();
    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };
    async function savePasaranOnline() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (select_pasaranonline == "" || select_pasaranonline == null) {
            flag = true;
            msg += "The Pasaran Online is required";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranonline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "PASARAN-SAVE",
                    idpasaran: idcomppasaran,
                    pasaran_hariraya: select_pasaranonline,
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;

                dispatch("handleRefreshPasaranOnline", idcomppasaran);
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function SaveTransaksi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_situs_field == "") {
            flag = true;
            msg += "The situs is required\n";
        }
        if (pasaran_diundi_field == "") {
            flag = true;
            msg += "The hari diundi is required\n";
        }
        if (pasaran_tutup_field == "") {
            flag = true;
            msg += "The pasaran tutup is required\n";
        }
        if (pasaran_jadwal_field == "") {
            flag = true;
            msg += "The pasaran jadwal is required\n";
        }
        if (pasaran_buka_field == "") {
            flag = true;
            msg += "The pasaran buka is required\n";
        }
        if (pasaran_display_field == "") {
            flag = true;
            msg += "The display is required\n";
        }

        if (flag == false) {
            const res = await fetch("/api/savepasaran", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "PASARAN-SAVE",
                    idpasaran: idcomppasaran,
                    pasaran_diundi: pasaran_diundi_field,
                    pasaran_url: pasaran_situs_field,
                    pasaran_jamtutup: pasaran_tutup_field,
                    pasaran_jamjadwal: pasaran_jadwal_field,
                    pasaran_jamopen: pasaran_buka_field,
                    pasaran_display: parseInt(pasaran_display_field),
                    pasaran_status: pasaran_status_field,
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function SaveLimitline() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_limitline4d_field == "") {
            flag = true;
            msg += "The Limitline 4D is required\n";
        }
        if (pasaran_limitline3d_field == "") {
            flag = true;
            msg += "The Limitline 3D is required\n";
        }
        if (pasaran_limitline3dd_field == "") {
            flag = true;
            msg += "The Limitline 3DD is required\n";
        }
        if (pasaran_limitline2d_field == "") {
            flag = true;
            msg += "The Limitline 2D is required\n";
        }
        if (pasaran_limitline2dd_field == "") {
            flag = true;
            msg += "The Limitline 2DD is required\n";
        }
        if (pasaran_limitline2dt_field == "") {
            flag = true;
            msg += "The Limitline 2DT is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranlimitline", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    pasaran_limitline4d: parseInt(pasaran_limitline4d_field),
                    pasaran_limitline3d: parseInt(pasaran_limitline3d_field),
                    pasaran_limitline3dd: parseInt(pasaran_limitline3dd_field),
                    pasaran_limitline2d: parseInt(pasaran_limitline2d_field),
                    pasaran_limitline2dd: parseInt(pasaran_limitline2dd_field),
                    pasaran_limitline2dt: parseInt(pasaran_limitline2dt_field),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function save432d() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_432d_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet4d_432d_field == "") {
            flag = true;
            msg += "The Max Bet 4D is required\n";
        }
        if (pasaran_maxbet3d_432d_field == "") {
            flag = true;
            msg += "The Max Bet 3D is required\n";
        }
        if (pasaran_maxbet3dd_432d_field == "") {
            flag = true;
            msg += "The Max Bet 3DD is required\n";
        }
        if (pasaran_maxbet2d_432d_field == "") {
            flag = true;
            msg += "The Max Bet 2D is required\n";
        }
        if (pasaran_maxbet2dd_432d_field == "") {
            flag = true;
            msg += "The Max Bet 2DD is required\n";
        }
        if (pasaran_maxbet2dt_432d_field == "") {
            flag = true;
            msg += "The Max Bet 2DT is required\n";
        }
        if (pasaran_limitotal4d_432d_field == "") {
            flag = true;
            msg += "The Limit Total 4D is required\n";
        }
        if (pasaran_limitotal3d_432d_field == "") {
            flag = true;
            msg += "The Limit Total 3D is required\n";
        }
        if (pasaran_limitotal3dd_432d_field == "") {
            flag = true;
            msg += "The Limit Total 3DD is required\n";
        }
        if (pasaran_limitotal2d_432d_field == "") {
            flag = true;
            msg += "The Limit Total 2D is required\n";
        }
        if (pasaran_limitotal2dd_432d_field == "") {
            flag = true;
            msg += "The Limit Total 2DD is required\n";
        }
        if (pasaran_limitotal2dt_432d_field == "") {
            flag = true;
            msg += "The Limit Total 2DT is required\n";
        }
        if (pasaran_limitglobal4d_432d_field == "") {
            flag = true;
            msg += "The Limit Global 4D is required\n";
        }
        if (pasaran_limitglobal3d_432d_field == "") {
            flag = true;
            msg += "The Limit Global 3D is required\n";
        }
        if (pasaran_limitglobal3dd_432d_field == "") {
            flag = true;
            msg += "The Limit Global 3DD is required\n";
        }
        if (pasaran_limitglobal2d_432d_field == "") {
            flag = true;
            msg += "The Limit Global 3D is required\n";
        }
        if (pasaran_limitglobal2dd_432d_field == "") {
            flag = true;
            msg += "The Limit Global 2DD is required\n";
        }
        if (pasaran_limitglobal2dt_432d_field == "") {
            flag = true;
            msg += "The Limit Global 2DT is required\n";
        }
        if (pasaran_win4d_432d_field == "") {
            flag = true;
            msg += "The Win 4D is required\n";
        }
        if (pasaran_win3d_432d_field == "") {
            flag = true;
            msg += "The Win 3D is required\n";
        }
        if (pasaran_win3dd_432d_field == "") {
            flag = true;
            msg += "The Win 3DD is required\n";
        }
        if (pasaran_win2d_432d_field == "") {
            flag = true;
            msg += "The Win 2D is required\n";
        }
        if (pasaran_win2dd_432d_field == "") {
            flag = true;
            msg += "The Win 2DD is required\n";
        }
        if (pasaran_win2dt_432d_field == "") {
            flag = true;
            msg += "The Win 2DT is required\n";
        }
        if (pasaran_disc4d_432d_field == "") {
            flag = true;
            msg += "The Disc 4D is required\n";
        }
        if (pasaran_disc3d_432d_field == "") {
            flag = true;
            msg += "The Disc 3D is required\n";
        }
        if (pasaran_disc3dd_432d_field == "") {
            flag = true;
            msg += "The Disc 3DD is required\n";
        }
        if (pasaran_disc2d_432d_field == "") {
            flag = true;
            msg += "The Disc 2D is required\n";
        }
        if (pasaran_disc2dd_432d_field == "") {
            flag = true;
            msg += "The Disc 2DD is required\n";
        }
        if (pasaran_disc2dt_432d_field == "") {
            flag = true;
            msg += "The Disc 2DT is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconf432", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_432d: parseInt(pasaran_minbet_432d_field),
                    pasaran_maxbet4d_432d: parseInt(pasaran_maxbet4d_432d_field),
                    pasaran_maxbet3d_432d: parseInt(pasaran_maxbet3d_432d_field),
                    pasaran_maxbet3dd_432d: parseInt(pasaran_maxbet3dd_432d_field),
                    pasaran_maxbet2d_432d: parseInt(pasaran_maxbet2d_432d_field),
                    pasaran_maxbet2dd_432d: parseInt(pasaran_maxbet2dd_432d_field),
                    pasaran_maxbet2dt_432d: parseInt(pasaran_maxbet2dt_432d_field),
                    pasaran_limitotal4d_432d: parseInt(pasaran_limitotal4d_432d_field),
                    pasaran_limitotal3d_432d: parseInt(pasaran_limitotal3d_432d_field),
                    pasaran_limitotal3dd_432d: parseInt(pasaran_limitotal3dd_432d_field),
                    pasaran_limitotal2d_432d: parseInt(pasaran_limitotal2d_432d_field),
                    pasaran_limitotal2dd_432d: parseInt(pasaran_limitotal2dd_432d_field),
                    pasaran_limitotal2dt_432d: parseInt(pasaran_limitotal2dt_432d_field),
                    pasaran_limitglobal4d_432d: parseInt(pasaran_limitglobal4d_432d_field),
                    pasaran_limitglobal3d_432d: parseInt(pasaran_limitglobal3d_432d_field),
                    pasaran_limitglobal3dd_432d: parseInt(pasaran_limitglobal3dd_432d_field),
                    pasaran_limitglobal2d_432d: parseInt(pasaran_limitglobal2d_432d_field),
                    pasaran_limitglobal2dd_432d: parseInt(pasaran_limitglobal2dd_432d_field),
                    pasaran_limitglobal2dt_432d: parseInt(pasaran_limitglobal2dt_432d_field),
                    pasaran_win4d_432d: parseInt(pasaran_win4d_432d_field),
                    pasaran_win3d_432d: parseInt(pasaran_win3d_432d_field),
                    pasaran_win3dd_432d: parseInt(pasaran_win3dd_432d_field),
                    pasaran_win2d_432d: parseInt(pasaran_win2d_432d_field),
                    pasaran_win2dd_432d: parseInt(pasaran_win2dd_432d_field),
                    pasaran_win2dt_432d: parseInt(pasaran_win2dt_432d_field),
                    pasaran_disc4d_432d: parseFloat(pasaran_disc4d_432d_field / 100),
                    pasaran_disc3d_432d: parseFloat(pasaran_disc3d_432d_field / 100),
                    pasaran_disc3dd_432d: parseFloat(pasaran_disc3dd_432d_field / 100),
                    pasaran_disc2d_432d: parseFloat(
                        pasaran_disc2d_432d_field / 100
                    ),
                    pasaran_disc2dd_432d: parseFloat(
                        pasaran_disc2dd_432d_field / 100
                    ),
                    pasaran_disc2dt_432d: parseFloat(
                        pasaran_disc2dt_432d_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveColokBebas() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_cbebas_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_cbebas_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limitotal_cbebas_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_cbebas_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_win_cbebas_field == "") {
            flag = true;
            msg += "The Win is required\n";
        }
        if (pasaran_disc_cbebas_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfcolokbebas", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    page: "PASARAN-SAVE",
                    idpasarantogel: pasaran_idpasarantogel_field,
                    pasaran_minbet_cbebas: parseInt(
                        pasaran_minbet_cbebas_field
                    ),
                    pasaran_maxbet_cbebas: parseInt(
                        pasaran_maxbet_cbebas_field
                    ),
                    pasaran_limitotal_cbebas: parseInt(
                        pasaran_limitotal_cbebas_field
                    ),
                    pasaran_limitglobal_cbebas: parseInt(
                        pasaran_limitglobal_cbebas_field
                    ),
                    pasaran_win_cbebas: parseFloat(pasaran_win_cbebas_field),
                    pasaran_disc_cbebas: parseFloat(
                        pasaran_disc_cbebas_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveColokMacau() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_cmacau_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_cmacau_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limitotal_cmacau_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_cmacau_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_win2_cmacau_field == "") {
            flag = true;
            msg += "The Win 2 Digit is required\n";
        }
        if (pasaran_win3_cmacau_field == "") {
            flag = true;
            msg += "The Win 3 Digit is required\n";
        }
        if (pasaran_win4_cmacau_field == "") {
            flag = true;
            msg += "The Win 4 Digit is required\n";
        }
        if (pasaran_disc_cmacau_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfcolokmacau", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_cmacau: parseInt(
                        pasaran_minbet_cmacau_field
                    ),
                    pasaran_maxbet_cmacau: parseInt(
                        pasaran_maxbet_cmacau_field
                    ),
                    pasaran_limitotal_cmacau: parseInt(
                        pasaran_limitotal_cmacau_field
                    ),
                    pasaran_limitglobal_cmacau: parseInt(
                        pasaran_limitglobal_cmacau_field
                    ),
                    pasaran_win2_cmacau: parseFloat(pasaran_win2_cmacau_field),
                    pasaran_win3_cmacau: parseFloat(pasaran_win3_cmacau_field),
                    pasaran_win4_cmacau: parseFloat(pasaran_win4_cmacau_field),
                    pasaran_disc_cmacau: parseFloat(
                        pasaran_disc_cmacau_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveColokNaga() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_cnaga_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_cnaga_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_cnaga_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_cnaga_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_win3_cnaga_field == "") {
            flag = true;
            msg += "The Win 3 Digit is required\n";
        }
        if (pasaran_win4_cnaga_field == "") {
            flag = true;
            msg += "The Win 4 Digit is required\n";
        }
        if (pasaran_disc_cnaga_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfcoloknaga", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_cnaga: parseInt(pasaran_minbet_cnaga_field),
                    pasaran_maxbet_cnaga: parseInt(pasaran_maxbet_cnaga_field),
                    pasaran_limittotal_cnaga: parseInt(
                        pasaran_limittotal_cnaga_field
                    ),
                    pasaran_limitglobal_cnaga: parseInt(
                        pasaran_limitglobal_cnaga_field
                    ),
                    pasaran_win3_cnaga: parseFloat(pasaran_win3_cnaga_field),
                    pasaran_win4_cnaga: parseFloat(pasaran_win4_cnaga_field),
                    pasaran_disc_cnaga: parseFloat(
                        pasaran_disc_cnaga_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveColokJitu() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_cjitu_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_cjitu_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_cjitu_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_cjitu_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_winas_cjitu_field == "") {
            flag = true;
            msg += "The Win AS is required\n";
        }
        if (pasaran_winkop_cjitu_field == "") {
            flag = true;
            msg += "The Win KOP is required\n";
        }
        if (pasaran_winkepala_cjitu_field == "") {
            flag = true;
            msg += "The Win KEPALA is required\n";
        }
        if (pasaran_winekor_cjitu_field == "") {
            flag = true;
            msg += "The Win EKOR is required\n";
        }
        if (pasaran_desc_cjitu_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfcolokjitu", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_cjitu: parseInt(pasaran_minbet_cjitu_field),
                    pasaran_maxbet_cjitu: parseInt(pasaran_maxbet_cjitu_field),
                    pasaran_limittotal_cjitu: parseInt(
                        pasaran_limittotal_cjitu_field
                    ),
                    pasaran_limitglobal_cjitu: parseInt(
                        pasaran_limitglobal_cjitu_field
                    ),
                    pasaran_winas_cjitu: parseFloat(pasaran_winas_cjitu_field),
                    pasaran_winkop_cjitu: parseFloat(
                        pasaran_winkop_cjitu_field
                    ),
                    pasaran_winkepala_cjitu: parseFloat(
                        pasaran_winkepala_cjitu_field
                    ),
                    pasaran_winekor_cjitu: parseFloat(
                        pasaran_winekor_cjitu_field
                    ),
                    pasaran_desc_cjitu: parseFloat(
                        pasaran_desc_cjitu_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function save5050Umum() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_5050umum_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_5050umum_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_5050umum_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_5050umum_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_keibesar_5050umum_field == "") {
            flag = true;
            msg += "The Kei Besar is required\n";
        }
        if (pasaran_keikecil_5050umum_field == "") {
            flag = true;
            msg += "The Kei Kecil is required\n";
        }
        if (pasaran_keigenap_5050umum_field == "") {
            flag = true;
            msg += "The Kei Genap is required\n";
        }
        if (pasaran_keiganjil_5050umum_field == "") {
            flag = true;
            msg += "The Kei Ganjil is required\n";
        }
        if (pasaran_keitengah_5050umum_field == "") {
            flag = true;
            msg += "The Kei Tengah is required\n";
        }
        if (pasaran_keitepi_5050umum_field == "") {
            flag = true;
            msg += "The Kei Tepi is required\n";
        }
        if (pasaran_discbesar_5050umum_field == "") {
            flag = true;
            msg += "The Disc Besar is required\n";
        }
        if (pasaran_disckecil_5050umum_field == "") {
            flag = true;
            msg += "The Disc Kecil is required\n";
        }
        if (pasaran_discgenap_5050umum_field == "") {
            flag = true;
            msg += "The Disc Genap is required\n";
        }
        if (pasaran_discganjil_5050umum_field == "") {
            flag = true;
            msg += "The Disc Ganjil is required\n";
        }
        if (pasaran_disctengah_5050umum_field == "") {
            flag = true;
            msg += "The Disc Tengah is required\n";
        }
        if (pasaran_disctepi_5050umum_field == "") {
            flag = true;
            msg += "The Disc Tepi is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconf5050umum", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_5050umum: parseInt(
                        pasaran_minbet_5050umum_field
                    ),
                    pasaran_maxbet_5050umum: parseInt(
                        pasaran_maxbet_5050umum_field
                    ),
                    pasaran_limittotal_5050umum: parseInt(
                        pasaran_limittotal_5050umum_field
                    ),
                    pasaran_limitglobal_5050umum: parseInt(
                        pasaran_limitglobal_5050umum_field
                    ),
                    pasaran_keibesar_5050umum: parseFloat(
                        (pasaran_keibesar_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_keikecil_5050umum: parseFloat(
                        (pasaran_keikecil_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_keigenap_5050umum: parseFloat(
                        (pasaran_keigenap_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_keiganjil_5050umum: parseFloat(
                        (pasaran_keiganjil_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_keitengah_5050umum: parseFloat(
                        (pasaran_keitengah_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_keitepi_5050umum: parseFloat(
                        (pasaran_keitepi_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_discbesar_5050umum: parseFloat(
                        (pasaran_discbesar_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_disckecil_5050umum: parseFloat(
                        (pasaran_disckecil_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_discgenap_5050umum: parseFloat(
                        (pasaran_discgenap_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_discganjil_5050umum: parseFloat(
                        (pasaran_discganjil_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_disctengah_5050umum: parseFloat(
                        (pasaran_disctengah_5050umum_field / 100).toPrecision(3)
                    ),
                    pasaran_disctepi_5050umum: parseFloat(
                        (pasaran_disctepi_5050umum_field / 100).toPrecision(3)
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function save5050Special() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_5050special_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_5050special_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_5050special_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_5050special_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_keiasganjil_5050special_field == "") {
            flag = true;
            msg += "The Kei AS Ganjil is required\n";
        }
        if (pasaran_keiasgenap_5050special_field == "") {
            flag = true;
            msg += "The Kei As Genap is required\n";
        }
        if (pasaran_keiasbesar_5050special_field == "") {
            flag = true;
            msg += "The Kei As Besar is required\n";
        }
        if (pasaran_keiaskecil_5050special_field == "") {
            flag = true;
            msg += "The Kei As Kecil is required\n";
        }
        if (pasaran_keikopganjil_5050special_field == "") {
            flag = true;
            msg += "The Kei Kop Ganjil is required\n";
        }
        if (pasaran_keikopgenap_5050special_field == "") {
            flag = true;
            msg += "The Kei Kop Genap is required\n";
        }
        if (pasaran_keikopbesar_5050special_field == "") {
            flag = true;
            msg += "The Kei Kop Besar is required\n";
        }
        if (pasaran_keikopkecil_5050special_field == "") {
            flag = true;
            msg += "The Kei Kop Kecil is required\n";
        }
        if (pasaran_keikepalaganjil_5050special_field == "") {
            flag = true;
            msg += "The Kei Kepala Ganjil is required\n";
        }
        if (pasaran_keikepalagenap_5050special_field == "") {
            flag = true;
            msg += "The Kei Kepala Genap is required\n";
        }
        if (pasaran_keikepalabesar_5050special_field == "") {
            flag = true;
            msg += "The Kei Kepala Besar is required\n";
        }
        if (pasaran_keikepalakecil_5050special_field == "") {
            flag = true;
            msg += "The Kei Kepala Kecil is required\n";
        }
        if (pasaran_keiekorganjil_5050special_field == "") {
            flag = true;
            msg += "The Kei Ekor Ganjil is required\n";
        }
        if (pasaran_keiekorgenap_5050special_field == "") {
            flag = true;
            msg += "The Kei Ekor Genap is required\n";
        }
        if (pasaran_keiekorbesar_5050special_field == "") {
            flag = true;
            msg += "The Kei Ekor Besar is required\n";
        }
        if (pasaran_keiekorkecil_5050special_field == "") {
            flag = true;
            msg += "The Kei Ekor Kecil is required\n";
        }
        if (pasaran_discasganjil_5050special_field == "") {
            flag = true;
            msg += "The Disc As Ganjil is required\n";
        }
        if (pasaran_discasgenap_5050special_field == "") {
            flag = true;
            msg += "The Disc As Genap is required\n";
        }
        if (pasaran_discasbesar_5050special_field == "") {
            flag = true;
            msg += "The Disc As Besar is required\n";
        }
        if (pasaran_discaskecil_5050special_field == "") {
            flag = true;
            msg += "The Disc As Kecil is required\n";
        }
        if (pasaran_disckopganjil_5050special_field == "") {
            flag = true;
            msg += "The Disc Kop Ganjil is required\n";
        }
        if (pasaran_disckopgenap_5050special_field == "") {
            flag = true;
            msg += "The Disc Kop Genap is required\n";
        }
        if (pasaran_disckopbesar_5050special_field == "") {
            flag = true;
            msg += "The Disc Kop Besar is required\n";
        }
        if (pasaran_disckopkecil_5050special_field == "") {
            flag = true;
            msg += "The Disc Kop Kecil is required\n";
        }
        if (pasaran_disckepalaganjil_5050special_field == "") {
            flag = true;
            msg += "The Disc Kepala Ganjil is required\n";
        }
        if (pasaran_disckepalagenap_5050special_field == "") {
            flag = true;
            msg += "The Disc Kepala Genap is required\n";
        }
        if (pasaran_disckepalabesar_5050special_field == "") {
            flag = true;
            msg += "The Disc Kepala Besar is required\n";
        }
        if (pasaran_disckepalakecil_5050special_field == "") {
            flag = true;
            msg += "The Disc Kepala Kecil is required\n";
        }
        if (pasaran_discekorganjil_5050special_field == "") {
            flag = true;
            msg += "The Disc Ekor Ganjil is required\n";
        }
        if (pasaran_discekorgenap_5050special_field == "") {
            flag = true;
            msg += "The Disc Ekor Genap is required\n";
        }
        if (pasaran_discekorbesar_5050special_field == "") {
            flag = true;
            msg += "The Disc Ekor Besar is required\n";
        }
        if (pasaran_discekorkecil_5050special_field == "") {
            flag = true;
            msg += "The Disc Ekor Kecil is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconf5050special", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    Pasaran_minbet_5050special: parseInt(
                        pasaran_minbet_5050special_field
                    ),
                    Pasaran_maxbet_5050special: parseInt(
                        pasaran_maxbet_5050special_field
                    ),
                    Pasaran_limitglobal_5050special: parseInt(
                        pasaran_limitglobal_5050special_field
                    ),
                    Pasaran_limittotal_5050special: parseInt(
                        pasaran_limittotal_5050special_field
                    ),
                    Pasaran_keiasganjil_5050special: parseFloat(
                        pasaran_keiasganjil_5050special_field / 100
                    ),
                    Pasaran_keiasgenap_5050special: parseFloat(
                        pasaran_keiasgenap_5050special_field / 100
                    ),
                    Pasaran_keiasbesar_5050special: parseFloat(
                        pasaran_keiasbesar_5050special_field / 100
                    ),
                    Pasaran_keiaskecil_5050special: parseFloat(
                        pasaran_keiaskecil_5050special_field / 100
                    ),
                    Pasaran_keikopganjil_5050special: parseFloat(
                        pasaran_keikopganjil_5050special_field / 100
                    ),
                    Pasaran_keikopgenap_5050special: parseFloat(
                        pasaran_keikopgenap_5050special_field / 100
                    ),
                    Pasaran_keikopbesar_5050special: parseFloat(
                        pasaran_keikopbesar_5050special_field / 100
                    ),
                    Pasaran_keikopkecil_5050special: parseFloat(
                        pasaran_keikopkecil_5050special_field / 100
                    ),
                    Pasaran_keikepalaganjil_5050special: parseFloat(
                        pasaran_keikepalaganjil_5050special_field / 100
                    ),
                    Pasaran_keikepalagenap_5050special: parseFloat(
                        pasaran_keikepalagenap_5050special_field / 100
                    ),
                    Pasaran_keikepalabesar_5050special: parseFloat(
                        pasaran_keikepalabesar_5050special_field / 100
                    ),
                    Pasaran_keikepalakecil_5050special: parseFloat(
                        pasaran_keikepalakecil_5050special_field / 100
                    ),
                    Pasaran_keiekorganjil_5050special: parseFloat(
                        pasaran_keiekorganjil_5050special_field / 100
                    ),
                    Pasaran_keiekorgenap_5050special: parseFloat(
                        pasaran_keiekorgenap_5050special_field / 100
                    ),
                    Pasaran_keiekorbesar_5050special: parseFloat(
                        pasaran_keiekorbesar_5050special_field / 100
                    ),
                    Pasaran_keiekorkecil_5050special: parseFloat(
                        pasaran_keiekorkecil_5050special_field / 100
                    ),
                    Pasaran_discasganjil_5050special: parseFloat(
                        pasaran_discasganjil_5050special_field / 100
                    ),
                    Pasaran_discasgenap_5050special: parseFloat(
                        pasaran_discasgenap_5050special_field / 100
                    ),
                    Pasaran_discasbesar_5050special: parseFloat(
                        pasaran_discasbesar_5050special_field / 100
                    ),
                    Pasaran_discaskecil_5050special: parseFloat(
                        pasaran_discaskecil_5050special_field / 100
                    ),
                    Pasaran_disckopganjil_5050special: parseFloat(
                        pasaran_disckopganjil_5050special_field / 100
                    ),
                    Pasaran_disckopgenap_5050special: parseFloat(
                        pasaran_disckopgenap_5050special_field / 100
                    ),
                    Pasaran_disckopbesar_5050special: parseFloat(
                        pasaran_disckopbesar_5050special_field / 100
                    ),
                    Pasaran_disckopkecil_5050special: parseFloat(
                        pasaran_disckopkecil_5050special_field / 100
                    ),
                    Pasaran_disckepalaganjil_5050special: parseFloat(
                        pasaran_disckepalaganjil_5050special_field / 100
                    ),
                    Pasaran_disckepalagenap_5050special: parseFloat(
                        pasaran_disckepalagenap_5050special_field / 100
                    ),
                    Pasaran_disckepalabesar_5050special: parseFloat(
                        pasaran_disckepalabesar_5050special_field / 100
                    ),
                    Pasaran_disckepalakecil_5050special: parseFloat(
                        pasaran_disckepalakecil_5050special_field / 100
                    ),
                    Pasaran_discekorganjil_5050special: parseFloat(
                        pasaran_discekorganjil_5050special_field / 100
                    ),
                    Pasaran_discekorgenap_5050special: parseFloat(
                        pasaran_discekorgenap_5050special_field / 100
                    ),
                    Pasaran_discekorbesar_5050special: parseFloat(
                        pasaran_discekorbesar_5050special_field / 100
                    ),
                    Pasaran_discekorkecil_5050special: parseFloat(
                        pasaran_discekorkecil_5050special_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function save5050Kombinasi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_5050kombinasi_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_5050kombinasi_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_5050kombinasi_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_5050kombinasi_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_belakangkeimono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Kei Mono is required\n";
        }
        if (pasaran_belakangkeistereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Kei Stereo is required\n";
        }
        if (pasaran_belakangkeikembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Kei Kembang is required\n";
        }
        if (pasaran_belakangkeikempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Kei Kempis is required\n";
        }
        if (pasaran_belakangkeikembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Kei Kembar is required\n";
        }
        if (pasaran_tengahkeimono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Kei Mono is required\n";
        }
        if (pasaran_tengahkeistereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Kei Stereo is required\n";
        }
        if (pasaran_tengahkeikembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Kei Kembang is required\n";
        }
        if (pasaran_tengahkeikempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Kei Kempis is required\n";
        }
        if (pasaran_tengahkeikembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Kei Kembar is required\n";
        }
        if (pasaran_depankeimono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Kei Mono is required\n";
        }
        if (pasaran_depankeistereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Kei Stereo is required\n";
        }
        if (pasaran_depankeikembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Kei Kembang is required\n";
        }
        if (pasaran_depankeikempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Kei Kempis is required\n";
        }
        if (pasaran_depankeikembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Kei Kembar is required\n";
        }
        if (pasaran_belakangdiscmono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Disc Mono is required\n";
        }
        if (pasaran_belakangdiscstereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Disc Stereo is required\n";
        }
        if (pasaran_belakangdisckembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Disc Kembang is required\n";
        }
        if (pasaran_belakangdisckempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Disc Kempis is required\n";
        }
        if (pasaran_belakangdisckembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Belakang Disc Kembar is required\n";
        }
        if (pasaran_tengahdiscmono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Disc Mono is required\n";
        }
        if (pasaran_tengahdiscstereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Disc Stereo is required\n";
        }
        if (pasaran_tengahdisckembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Disc Kembang is required\n";
        }
        if (pasaran_tengahdisckempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Disc Kempis is required\n";
        }
        if (pasaran_tengahdisckembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Tengah Disc Kembar is required\n";
        }
        if (pasaran_depandiscmono_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Disc Mono is required\n";
        }
        if (pasaran_depandiscstereo_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Disc Stereo is required\n";
        }
        if (pasaran_depandisckembang_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Disc Kembang is required\n";
        }
        if (pasaran_depandisckempis_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Disc Kempis is required\n";
        }
        if (pasaran_depandisckembar_5050kombinasi_field == "") {
            flag = true;
            msg += "The Depan Disc Kembar is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconf5050kombinasi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_5050kombinasi: parseInt(
                        pasaran_minbet_5050kombinasi_field
                    ),
                    pasaran_maxbet_5050kombinasi: parseInt(
                        pasaran_maxbet_5050kombinasi_field
                    ),
                    pasaran_limitglobal_5050kombinasi: parseInt(
                        pasaran_limitglobal_5050kombinasi_field
                    ),
                    pasaran_limittotal_5050kombinasi: parseInt(
                        pasaran_limittotal_5050kombinasi_field
                    ),
                    pasaran_belakangkeimono_5050kombinasi: parseFloat(
                        pasaran_belakangkeimono_5050kombinasi_field / 100
                    ),
                    pasaran_belakangkeistereo_5050kombinasi: parseFloat(
                        pasaran_belakangkeistereo_5050kombinasi_field / 100
                    ),
                    pasaran_belakangkeikembang_5050kombinasi: parseFloat(
                        pasaran_belakangkeikembang_5050kombinasi_field / 100
                    ),
                    pasaran_belakangkeikempis_5050kombinasi: parseFloat(
                        pasaran_belakangkeikempis_5050kombinasi_field / 100
                    ),
                    pasaran_belakangkeikembar_5050kombinasi: parseFloat(
                        pasaran_belakangkeikembar_5050kombinasi_field / 100
                    ),
                    pasaran_tengahkeimono_5050kombinasi: parseFloat(
                        pasaran_tengahkeimono_5050kombinasi_field / 100
                    ),
                    pasaran_tengahkeistereo_5050kombinasi: parseFloat(
                        pasaran_tengahkeistereo_5050kombinasi_field / 100
                    ),
                    pasaran_tengahkeikembang_5050kombinasi: parseFloat(
                        pasaran_tengahkeikembang_5050kombinasi_field / 100
                    ),
                    pasaran_tengahkeikempis_5050kombinasi: parseFloat(
                        pasaran_tengahkeikempis_5050kombinasi_field / 100
                    ),
                    pasaran_tengahkeikembar_5050kombinasi: parseFloat(
                        pasaran_tengahkeikembar_5050kombinasi_field / 100
                    ),
                    pasaran_depankeimono_5050kombinasi: parseFloat(
                        pasaran_depankeimono_5050kombinasi_field / 100
                    ),
                    pasaran_depankeistereo_5050kombinasi: parseFloat(
                        pasaran_depankeistereo_5050kombinasi_field / 100
                    ),
                    pasaran_depankeikembang_5050kombinasi: parseFloat(
                        pasaran_depankeikembang_5050kombinasi_field / 100
                    ),
                    pasaran_depankeikempis_5050kombinasi: parseFloat(
                        pasaran_depankeikempis_5050kombinasi_field / 100
                    ),
                    pasaran_depankeikembar_5050kombinasi: parseFloat(
                        pasaran_depankeikembar_5050kombinasi_field / 100
                    ),
                    pasaran_belakangdiscmono_5050kombinasi: parseFloat(
                        pasaran_belakangdiscmono_5050kombinasi_field / 100
                    ),
                    pasaran_belakangdiscstereo_5050kombinasi: parseFloat(
                        pasaran_belakangdiscstereo_5050kombinasi_field / 100
                    ),
                    pasaran_belakangdisckembang_5050kombinasi: parseFloat(
                        pasaran_belakangdisckembang_5050kombinasi_field / 100
                    ),
                    pasaran_belakangdisckempis_5050kombinasi: parseFloat(
                        pasaran_belakangdisckempis_5050kombinasi_field / 100
                    ),
                    pasaran_belakangdisckembar_5050kombinasi: parseFloat(
                        pasaran_belakangdisckembar_5050kombinasi_field / 100
                    ),
                    pasaran_tengahdiscmono_5050kombinasi: parseFloat(
                        pasaran_tengahdiscmono_5050kombinasi_field / 100
                    ),
                    pasaran_tengahdiscstereo_5050kombinasi: parseFloat(
                        pasaran_tengahdiscstereo_5050kombinasi_field / 100
                    ),
                    pasaran_tengahdisckembang_5050kombinasi: parseFloat(
                        pasaran_tengahdisckembang_5050kombinasi_field / 100
                    ),
                    pasaran_tengahdisckempis_5050kombinasi: parseFloat(
                        pasaran_tengahdisckempis_5050kombinasi_field / 100
                    ),
                    pasaran_tengahdisckembar_5050kombinasi: parseFloat(
                        pasaran_tengahdisckembar_5050kombinasi_field / 100
                    ),
                    pasaran_depandiscmono_5050kombinasi: parseFloat(
                        pasaran_depandiscmono_5050kombinasi_field / 100
                    ),
                    pasaran_depandiscstereo_5050kombinasi: parseFloat(
                        pasaran_depandiscstereo_5050kombinasi_field / 100
                    ),
                    pasaran_depandisckembang_5050kombinasi: parseFloat(
                        pasaran_depandisckembang_5050kombinasi_field / 100
                    ),
                    pasaran_depandisckempis_5050kombinasi: parseFloat(
                        pasaran_depandisckempis_5050kombinasi_field / 100
                    ),
                    pasaran_depandisckembar_5050kombinasi: parseFloat(
                        pasaran_depandisckembar_5050kombinasi_field / 100
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveMacauKombinasi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_kombinasi_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_kombinasi_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_kombinasi_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_kombinasi_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_win_kombinasi_field == "") {
            flag = true;
            msg += "The Win is required\n";
        }
        if (pasaran_disc_kombinasi_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfmacaukombinasi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_kombinasi: parseInt(
                        pasaran_minbet_kombinasi_field
                    ),
                    pasaran_maxbet_kombinasi: parseInt(
                        pasaran_maxbet_kombinasi_field
                    ),
                    pasaran_limittotal_kombinasi: parseInt(
                        pasaran_limittotal_kombinasi_field
                    ),
                    pasaran_limitglobal_kombinasi: parseInt(
                        pasaran_limitglobal_kombinasi_field
                    ),
                    pasaran_win_kombinasi: parseFloat(
                        pasaran_win_kombinasi_field
                    ),
                    pasaran_disc_kombinasi: parseFloat(
                        (pasaran_disc_kombinasi_field / 100).toPrecision(3)
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveDasar() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_dasar_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_dasar_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_dasar_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_dasar_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_keibesar_dasar_field == "") {
            flag = true;
            msg += "The Kei Besar is required\n";
        }
        if (pasaran_keikecil_dasar_field == "") {
            flag = true;
            msg += "The Kei Kecil is required\n";
        }
        if (pasaran_keigenap_dasar_field == "") {
            flag = true;
            msg += "The Kei Genap is required\n";
        }
        if (pasaran_keiganjil_dasar_field == "") {
            flag = true;
            msg += "The Kei Ganjil is required\n";
        }
        if (pasaran_discbesar_dasar_field == "") {
            flag = true;
            msg += "The Diskon Besar is required\n";
        }
        if (pasaran_disckecil_dasar_field == "") {
            flag = true;
            msg += "The Diskon Kecil is required\n";
        }
        if (pasaran_discgenap_dasar_field == "") {
            flag = true;
            msg += "The Diskon Genap is required\n";
        }
        if (pasaran_discganjil_dasar_field == "") {
            flag = true;
            msg += "The Diskon Ganjil is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfdasar", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_dasar: parseInt(pasaran_minbet_dasar_field),
                    pasaran_maxbet_dasar: parseInt(pasaran_maxbet_dasar_field),
                    pasaran_limittotal_dasar: parseInt(
                        pasaran_limittotal_dasar_field
                    ),
                    pasaran_limitglobal_dasar: parseInt(
                        pasaran_limitglobal_dasar_field
                    ),
                    pasaran_keibesar_dasar: parseFloat(
                        (pasaran_keibesar_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_keikecil_dasar: parseFloat(
                        (pasaran_keikecil_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_keigenap_dasar: parseFloat(
                        (pasaran_keigenap_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_keiganjil_dasar: parseFloat(
                        (pasaran_keiganjil_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_discbesar_dasar: parseFloat(
                        (pasaran_discbesar_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_disckecil_dasar: parseFloat(
                        (pasaran_disckecil_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_discgenap_dasar: parseFloat(
                        (pasaran_discgenap_dasar_field / 100).toPrecision(3)
                    ),
                    pasaran_discganjil_dasar: parseFloat(
                        (pasaran_discganjil_dasar_field / 100).toPrecision(3)
                    ),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function saveShio() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (pasaran_minbet_shio_field == "") {
            flag = true;
            msg += "The Min Bet is required\n";
        }
        if (pasaran_maxbet_shio_field == "") {
            flag = true;
            msg += "The Max Bet is required\n";
        }
        if (pasaran_limittotal_shio_field == "") {
            flag = true;
            msg += "The Limit Total is required\n";
        }
        if (pasaran_limitglobal_shio_field == "") {
            flag = true;
            msg += "The Limit Global is required\n";
        }
        if (pasaran_win_shio_field == "") {
            flag = true;
            msg += "The Win is required\n";
        }

        if (pasaran_disc_shio_field == "") {
            flag = true;
            msg += "The Diskon is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasaranconfshio", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idpasaran: idcomppasaran,
                    idpasarantogel: pasaran_idpasarantogel_field,
                    page: "PASARAN-SAVE",
                    pasaran_minbet_shio: parseInt(pasaran_minbet_shio_field),
                    pasaran_maxbet_shio: parseInt(pasaran_maxbet_shio_field),
                    pasaran_limittotal_shio: parseInt(
                        pasaran_limittotal_shio_field
                    ),
                    pasaran_limitglobal_shio: parseInt(
                        pasaran_limitglobal_shio_field
                    ),
                    pasaran_shioyear_shio: pasaran_shioyear_shio_field,
                    pasaran_disc_shio: parseFloat(
                        (pasaran_disc_shio_field / 100).toPrecision(3)
                    ),
                    pasaran_win_shio: parseFloat(pasaran_win_shio_field),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    async function newOnline() {
        let myModal = new bootstrap.Modal(
            document.getElementById("modaNewOnline")
        );
        myModal.show();
    }
    const handleKeyboard_number = (e) => {
        let numbera;
        for (let i = 0; i < pasaran_display_field.length; i++) {
            numbera = parseInt(pasaran_display_field[i]);
            if (isNaN(numbera)) {
                pasaran_display_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline4d_field.length; i++) {
            numbera = parseInt(pasaran_limitline4d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline4d_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline3d_field.length; i++) {
            numbera = parseInt(pasaran_limitline3d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline3d_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline3dd_field.length; i++) {
            numbera = parseInt(pasaran_limitline3dd_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline3dd_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline2d_field.length; i++) {
            numbera = parseInt(pasaran_limitline2d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline2d_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline2dd_field.length; i++) {
            numbera = parseInt(pasaran_limitline2dd_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline2dd_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitline2dt_field.length; i++) {
            numbera = parseInt(pasaran_limitline2dt_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitline2dt_field = "";
            }
        }
        for (let i = 0; i < pasaran_minbet_432d_field.length; i++) {
            numbera = parseInt(pasaran_minbet_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet4d_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet4d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet4d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet3d_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet3d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet3d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet3dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet3dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet3dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet2d_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet2d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet2d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet2dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet2dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet2dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet2dt_432d_field.length; i++) {
            numbera = parseInt(pasaran_maxbet2dt_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet2dt_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal4d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal4d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal4d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal3d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal3d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal3d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal3dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal3dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal3dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal2d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal2d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal2d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal2dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal2dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal2dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal2dt_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitotal2dt_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal2dt_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal4d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal4d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal4d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal3d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal3d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal3d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal3dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal3dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal3dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal2d_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal2d_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal2d_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal2dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal2dd_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal2dd_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal2dt_432d_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal2dt_432d_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal2dt_432d_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win4d_432d_field.length; i++) {
            numbera = parseInt(pasaran_win4d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win4d_432d_field[i] != ".") {
                    pasaran_win4d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win3d_432d_field.length; i++) {
            numbera = parseInt(pasaran_win3d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win3d_432d_field[i] != ".") {
                    pasaran_win3d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win3dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_win3dd_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win3dd_432d_field[i] != ".") {
                    pasaran_win3dd_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win2d_432d_field.length; i++) {
            numbera = parseInt(pasaran_win2d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win2d_432d_field[i] != ".") {
                    pasaran_win2d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win2dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_win2dd_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win2dd_432d_field[i] != ".") {
                    pasaran_win2dd_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win2dt_432d_field.length; i++) {
            numbera = parseInt(pasaran_win2dt_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win2dt_432d_field[i] != ".") {
                    pasaran_win2dt_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc4d_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc4d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc4d_432d_field[i] != ".") {
                    pasaran_disc4d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc3d_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc3d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc3d_432d_field[i] != ".") {
                    pasaran_disc3d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc3dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc3dd_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc3dd_432d_field[i] != ".") {
                    pasaran_disc3dd_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc2d_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc2d_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc2d_432d_field[i] != ".") {
                    pasaran_disc2d_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc2dd_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc2dd_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc2dd_432d_field[i] != ".") {
                    pasaran_disc2dd_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc2dt_432d_field.length; i++) {
            numbera = parseInt(pasaran_disc2dt_432d_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc2dt_432d_field[i] != ".") {
                    pasaran_disc2dt_432d_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_minbet_cbebas_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_cbebas_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_cbebas_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_cbebas_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_limitotal_cbebas_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal_cbebas_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_cbebas_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_cbebas_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_win_cbebas_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win_cbebas_field[i] != ".") {
                    pasaran_win_cbebas_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc_cbebas_field.length; i++) {
            numbera = parseInt(pasaran_disc_cbebas_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc_cbebas_field[i] != ".") {
                    pasaran_disc_cbebas_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_minbet_cmacau_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_cmacau_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_cmacau_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_cmacau_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitotal_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_limitotal_cmacau_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitotal_cmacau_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_cmacau_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_cmacau_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win2_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_win2_cmacau_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win2_cmacau_field[i] != ".") {
                    pasaran_win2_cmacau_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win3_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_win3_cmacau_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win3_cmacau_field[i] != ".") {
                    pasaran_win3_cmacau_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win4_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_win4_cmacau_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win4_cmacau_field[i] != ".") {
                    pasaran_win4_cmacau_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc_cmacau_field.length; i++) {
            numbera = parseInt(pasaran_disc_cmacau_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc_cmacau_field[i] != ".") {
                    pasaran_disc_cmacau_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_minbet_cnaga_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_cnaga_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_cnaga_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_cnaga_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_cnaga_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_cnaga_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_cnaga_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_cnaga_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win3_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_win3_cnaga_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win3_cnaga_field[i] != ".") {
                    pasaran_win3_cnaga_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_win4_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_win4_cnaga_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win4_cnaga_field[i] != ".") {
                    pasaran_win4_cnaga_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc_cnaga_field.length; i++) {
            numbera = parseInt(pasaran_disc_cnaga_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc_cnaga_field[i] != ".") {
                    pasaran_disc_cnaga_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_minbet_cjitu_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_cjitu_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_cjitu_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_cjitu_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_cjitu_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_cjitu_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_cjitu_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_cjitu_field = 0;
            }
        }
        for (let i = 0; i < pasaran_winas_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_winas_cjitu_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_winas_cjitu_field[i] != ".") {
                    pasaran_winas_cjitu_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_winkop_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_winkop_cjitu_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_winkop_cjitu_field[i] != ".") {
                    pasaran_winkop_cjitu_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_winkepala_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_winkepala_cjitu_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_winkepala_cjitu_field[i] != ".") {
                    pasaran_winkepala_cjitu_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_winekor_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_winekor_cjitu_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_winekor_cjitu_field[i] != ".") {
                    pasaran_winekor_cjitu_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_desc_cjitu_field.length; i++) {
            numbera = parseInt(pasaran_desc_cjitu_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_desc_cjitu_field[i] != ".") {
                    pasaran_desc_cjitu_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_minbet_5050umum_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_5050umum_field = "";
            }
        }
        for (let i = 0; i < pasaran_maxbet_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_5050umum_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_5050umum_field = "";
            }
        }
        for (let i = 0; i < pasaran_limitglobal_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_5050umum_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_5050umum_field = "";
            }
        }
        for (let i = 0; i < pasaran_limittotal_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_5050umum_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_5050umum_field = "";
            }
        }
        for (let i = 0; i < pasaran_keibesar_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keibesar_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keibesar_5050umum_field[i] != "." &&
                    pasaran_keibesar_5050umum_field[i] != "-"
                ) {
                    pasaran_keibesar_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keikecil_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keikecil_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikecil_5050umum_field[i] != "." &&
                    pasaran_keikecil_5050umum_field[i] != "-"
                ) {
                    pasaran_keikecil_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keigenap_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keigenap_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keigenap_5050umum_field[i] != "." &&
                    pasaran_keigenap_5050umum_field[i] != "-"
                ) {
                    pasaran_keigenap_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keiganjil_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keiganjil_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiganjil_5050umum_field[i] != "." &&
                    pasaran_keiganjil_5050umum_field[i] != "-"
                ) {
                    pasaran_keiganjil_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keitengah_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keitengah_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keitengah_5050umum_field[i] != "." &&
                    pasaran_keitengah_5050umum_field[i] != "-"
                ) {
                    pasaran_keitengah_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keitepi_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_keitepi_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keitepi_5050umum_field[i] != "." &&
                    pasaran_keitepi_5050umum_field[i] != "-"
                ) {
                    pasaran_keitepi_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discbesar_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_discbesar_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discbesar_5050umum_field[i] != ".") {
                    pasaran_discbesar_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disckecil_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_disckecil_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckecil_5050umum_field[i] != ".") {
                    pasaran_disckecil_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discgenap_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_discgenap_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discgenap_5050umum_field[i] != ".") {
                    pasaran_discgenap_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discganjil_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_discganjil_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discganjil_5050umum_field[i] != ".") {
                    pasaran_discganjil_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disctengah_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_disctengah_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disctengah_5050umum_field[i] != ".") {
                    pasaran_disctengah_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disctepi_5050umum_field.length; i++) {
            numbera = parseInt(pasaran_disctepi_5050umum_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disctepi_5050umum_field[i] != ".") {
                    pasaran_disctepi_5050umum_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_5050special_field.length; i++) {
            numbera = parseInt(pasaran_minbet_5050special_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_5050special_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_5050special_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_5050special_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_5050special_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_5050special_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_5050special_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_5050special_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_5050special_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_5050special_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_5050special_field = 0;
            }
        }
        for (let i = 0; i < pasaran_keiasganjil_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keiasganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiasganjil_5050special_field[i] != "." &&
                    pasaran_keiasganjil_5050special_field[i] != "-"
                ) {
                    pasaran_keiasganjil_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keiasgenap_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keiasgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiasgenap_5050special_field[i] != "." &&
                    pasaran_keiasgenap_5050special_field[i] != "-"
                ) {
                    pasaran_keiasgenap_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keiasbesar_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keiasbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiasbesar_5050special_field[i] != "." &&
                    pasaran_keiasbesar_5050special_field[i] != "-"
                ) {
                    pasaran_keiasbesar_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keiaskecil_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keiaskecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiaskecil_5050special_field[i] != "." &&
                    pasaran_keiaskecil_5050special_field[i] != "-"
                ) {
                    pasaran_keiaskecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keikopganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keikopganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikopganjil_5050special_field[i] != "." &&
                    pasaran_keikopganjil_5050special_field[i] != "-"
                ) {
                    pasaran_keikopganjil_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keikopgenap_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keikopgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikopgenap_5050special_field[i] != "." &&
                    pasaran_keikopgenap_5050special_field[i] != "-"
                ) {
                    pasaran_keikopgenap_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keikopbesar_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keikopbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikopbesar_5050special_field[i] != "." &&
                    pasaran_keikopbesar_5050special_field[i] != "-"
                ) {
                    pasaran_keikopbesar_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keikopkecil_5050special_field.length; i++) {
            numbera = parseInt(pasaran_keikopkecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikopkecil_5050special_field[i] != "." &&
                    pasaran_keikopkecil_5050special_field[i] != "-"
                ) {
                    pasaran_keikopkecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keikepalaganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keikepalaganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikepalaganjil_5050special_field[i] != "." &&
                    pasaran_keikepalaganjil_5050special_field[i] != "-"
                ) {
                    pasaran_keikepalaganjil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keikepalagenap_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keikepalagenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikepalagenap_5050special_field[i] != "." &&
                    pasaran_keikepalagenap_5050special_field[i] != "-"
                ) {
                    pasaran_keikepalagenap_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keikepalabesar_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keikepalabesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikepalabesar_5050special_field[i] != "." &&
                    pasaran_keikepalabesar_5050special_field[i] != "-"
                ) {
                    pasaran_keikepalabesar_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keikepalakecil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keikepalakecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikepalakecil_5050special_field[i] != "." &&
                    pasaran_keikepalakecil_5050special_field[i] != "-"
                ) {
                    pasaran_keikepalakecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keiekorganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keiekorganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiekorganjil_5050special_field[i] != "." &&
                    pasaran_keiekorganjil_5050special_field[i] != "-"
                ) {
                    pasaran_keiekorganjil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keiekorgenap_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keiekorgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiekorgenap_5050special_field[i] != "." &&
                    pasaran_keiekorgenap_5050special_field[i] != "-"
                ) {
                    pasaran_keiekorgenap_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keiekorbesar_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keiekorbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiekorbesar_5050special_field[i] != "." &&
                    pasaran_keiekorbesar_5050special_field[i] != "-"
                ) {
                    pasaran_keiekorbesar_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_keiekorkecil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_keiekorkecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiekorkecil_5050special_field[i] != "." &&
                    pasaran_keiekorkecil_5050special_field[i] != "-"
                ) {
                    pasaran_keiekorkecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_discasganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_discasganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discasganjil_5050special_field[i] != ".") {
                    pasaran_discasganjil_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discasgenap_5050special_field.length; i++) {
            numbera = parseInt(pasaran_discasgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discasgenap_5050special_field[i] != ".") {
                    pasaran_discasgenap_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discasbesar_5050special_field.length; i++) {
            numbera = parseInt(pasaran_discasbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discasbesar_5050special_field[i] != ".") {
                    pasaran_discasbesar_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discaskecil_5050special_field.length; i++) {
            numbera = parseInt(pasaran_discaskecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discaskecil_5050special_field[i] != ".") {
                    pasaran_discaskecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckopganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckopganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckopganjil_5050special_field[i] != ".") {
                    pasaran_disckopganjil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckopgenap_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckopgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckopgenap_5050special_field[i] != ".") {
                    pasaran_disckopgenap_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckopbesar_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckopbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckopbesar_5050special_field[i] != ".") {
                    pasaran_disckopbesar_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckopkecil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckopkecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckopkecil_5050special_field[i] != ".") {
                    pasaran_disckopkecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckepalaganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckepalaganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckepalaganjil_5050special_field[i] != ".") {
                    pasaran_disckepalaganjil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckepalagenap_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckepalagenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckepalagenap_5050special_field[i] != ".") {
                    pasaran_disckepalagenap_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckepalabesar_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckepalabesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckepalabesar_5050special_field[i] != ".") {
                    pasaran_disckepalabesar_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_disckepalakecil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_disckepalakecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckepalakecil_5050special_field[i] != ".") {
                    pasaran_disckepalakecil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_discekorganjil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_discekorganjil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discekorganjil_5050special_field[i] != ".") {
                    pasaran_discekorganjil_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_discekorgenap_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_discekorgenap_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discekorgenap_5050special_field[i] != ".") {
                    pasaran_discekorgenap_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_discekorbesar_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_discekorbesar_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discekorbesar_5050special_field[i] != ".") {
                    pasaran_discekorbesar_5050special_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_discekorkecil_5050special_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_discekorkecil_5050special_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discekorkecil_5050special_field[i] != ".") {
                    pasaran_discekorkecil_5050special_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_5050kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_minbet_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_5050kombinasi_field = "";
            }
        }
        for (let i = 0; i < pasaran_maxbet_5050kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_5050kombinasi_field = "";
            }
        }
        for (
            let i = 0;
            i < pasaran_limitglobal_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_limitglobal_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_5050kombinasi_field = "";
            }
        }
        for (
            let i = 0;
            i < pasaran_limittotal_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_limittotal_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_5050kombinasi_field = "";
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangkeimono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_belakangkeimono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_belakangkeimono_5050kombinasi_field[i] != "." &&
                    pasaran_belakangkeimono_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_belakangkeimono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangkeistereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangkeistereo_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (
                    pasaran_belakangkeistereo_5050kombinasi_field[i] != "." &&
                    pasaran_belakangkeistereo_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_belakangkeistereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangkeikembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangkeikembang_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (
                    pasaran_belakangkeikembang_5050kombinasi_field[i] != "." &&
                    pasaran_belakangkeikembang_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_belakangkeikembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangkeikempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangkeikempis_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (
                    pasaran_belakangkeikempis_5050kombinasi_field[i] != "." &&
                    pasaran_belakangkeikempis_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_belakangkeikempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangkeikembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangkeikembar_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (
                    pasaran_belakangkeikembar_5050kombinasi_field[i] != "." &&
                    pasaran_belakangkeikembar_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_belakangkeikembar_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahkeimono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahkeimono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_tengahkeimono_5050kombinasi_field[i] != "." &&
                    pasaran_tengahkeimono_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_tengahkeimono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahkeistereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahkeistereo_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_tengahkeistereo_5050kombinasi_field[i] != "." &&
                    pasaran_tengahkeistereo_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_tengahkeistereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahkeikembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahkeikembang_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_tengahkeikembang_5050kombinasi_field[i] != "." &&
                    pasaran_tengahkeikembang_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_tengahkeikembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahkeikempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahkeikempis_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_tengahkeikempis_5050kombinasi_field[i] != "." &&
                    pasaran_tengahkeikempis_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_tengahkeikempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahkeikembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahkeikembar_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_tengahkeikembar_5050kombinasi_field[i] != "." &&
                    pasaran_tengahkeikembar_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_tengahkeikembar_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depankeimono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depankeimono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_depankeimono_5050kombinasi_field[i] != "." &&
                    pasaran_depankeimono_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_depankeimono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depankeistereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depankeistereo_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_depankeistereo_5050kombinasi_field[i] != "." &&
                    pasaran_depankeistereo_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_depankeistereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depankeikembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depankeikembang_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_depankeikembang_5050kombinasi_field[i] != "." &&
                    pasaran_depankeikembang_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_depankeikembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depankeikempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depankeikempis_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_depankeikempis_5050kombinasi_field[i] != "." &&
                    pasaran_depankeikempis_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_depankeikempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depankeikembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depankeikembar_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_depankeikembar_5050kombinasi_field[i] != "." &&
                    pasaran_depankeikembar_5050kombinasi_field[i] != "-"
                ) {
                    pasaran_depankeikembar_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangdiscmono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_belakangdiscmono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_belakangdiscmono_5050kombinasi_field[i] != ".") {
                    pasaran_belakangdiscmono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangdiscstereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangdiscstereo_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (pasaran_belakangdiscstereo_5050kombinasi_field[i] != ".") {
                    pasaran_belakangdiscstereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangdisckembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangdisckembang_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (pasaran_belakangdisckembang_5050kombinasi_field[i] != ".") {
                    pasaran_belakangdisckembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangdisckempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangdisckempis_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (pasaran_belakangdisckempis_5050kombinasi_field[i] != ".") {
                    pasaran_belakangdisckempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_belakangdisckembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_belakangdisckembar_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (pasaran_belakangdisckembar_5050kombinasi_field[i] != ".") {
                    pasaran_belakangdisckembar_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahdiscmono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahdiscmono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_tengahdiscmono_5050kombinasi_field[i] != ".") {
                    pasaran_tengahdiscmono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahdiscstereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahdiscstereo_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_tengahdiscstereo_5050kombinasi_field[i] != ".") {
                    pasaran_tengahdiscstereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahdisckembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(
                pasaran_tengahdisckembang_5050kombinasi_field[i]
            );
            if (isNaN(numbera)) {
                if (pasaran_tengahdisckembang_5050kombinasi_field[i] != ".") {
                    pasaran_tengahdisckembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahdisckempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahdisckempis_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_tengahdisckempis_5050kombinasi_field[i] != ".") {
                    pasaran_tengahdisckempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_tengahdisckembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_tengahdisckembar_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_tengahdisckembar_5050kombinasi_field[i] != ".") {
                    pasaran_tengahdisckembar_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depandiscmono_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depandiscmono_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_depandiscmono_5050kombinasi_field[i] != ".") {
                    pasaran_depandiscmono_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depandiscstereo_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depandiscstereo_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_depandiscstereo_5050kombinasi_field[i] != ".") {
                    pasaran_depandiscstereo_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depandisckembang_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depandisckembang_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_depandisckembang_5050kombinasi_field[i] != ".") {
                    pasaran_depandisckembang_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depandisckempis_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depandisckempis_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_depandisckempis_5050kombinasi_field[i] != ".") {
                    pasaran_depandisckempis_5050kombinasi_field = 0;
                }
            }
        }
        for (
            let i = 0;
            i < pasaran_depandisckembar_5050kombinasi_field.length;
            i++
        ) {
            numbera = parseInt(pasaran_depandisckembar_5050kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_depandisckembar_5050kombinasi_field[i] != ".") {
                    pasaran_depandisckembar_5050kombinasi_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_minbet_kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_kombinasi_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_kombinasi_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_kombinasi_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_kombinasi_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_win_kombinasi_field[i]);
            if (isNaN(numbera)) {
                pasaran_win_kombinasi_field = 0;
            }
        }
        for (let i = 0; i < pasaran_disc_kombinasi_field.length; i++) {
            numbera = parseInt(pasaran_disc_kombinasi_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc_kombinasi_field[i] != ".") {
                    pasaran_disc_kombinasi_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_dasar_field.length; i++) {
            numbera = parseInt(pasaran_minbet_dasar_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_dasar_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_dasar_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_dasar_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_dasar_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_dasar_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_dasar_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_dasar_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_dasar_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_dasar_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_dasar_field = 0;
            }
        }
        for (let i = 0; i < pasaran_keibesar_dasar_field.length; i++) {
            numbera = parseInt(pasaran_keibesar_dasar_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keibesar_dasar_field[i] != "." &&
                    pasaran_keibesar_dasar_field[i] != "-"
                ) {
                    pasaran_keibesar_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keikecil_dasar_field.length; i++) {
            numbera = parseInt(pasaran_keikecil_dasar_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keikecil_dasar_field[i] != "." &&
                    pasaran_keikecil_dasar_field[i] != "-"
                ) {
                    pasaran_keikecil_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keigenap_dasar_field.length; i++) {
            numbera = parseInt(pasaran_keigenap_dasar_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keigenap_dasar_field[i] != "." &&
                    pasaran_keigenap_dasar_field[i] != "-"
                ) {
                    pasaran_keigenap_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_keiganjil_dasar_field.length; i++) {
            numbera = parseInt(pasaran_keiganjil_dasar_field[i]);
            if (isNaN(numbera)) {
                if (
                    pasaran_keiganjil_dasar_field[i] != "." &&
                    pasaran_keiganjil_dasar_field[i] != "-"
                ) {
                    pasaran_keiganjil_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discbesar_dasar_field.length; i++) {
            numbera = parseInt(pasaran_discbesar_dasar_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discbesar_dasar_field[i] != ".") {
                    pasaran_discbesar_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disckecil_dasar_field.length; i++) {
            numbera = parseInt(pasaran_disckecil_dasar_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disckecil_dasar_field[i] != ".") {
                    pasaran_disckecil_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discgenap_dasar_field.length; i++) {
            numbera = parseInt(pasaran_discgenap_dasar_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discgenap_dasar_field[i] != ".") {
                    pasaran_discgenap_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_discganjil_dasar_field.length; i++) {
            numbera = parseInt(pasaran_discganjil_dasar_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_discganjil_dasar_field[i] != ".") {
                    pasaran_discganjil_dasar_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_minbet_shio_field.length; i++) {
            numbera = parseInt(pasaran_minbet_shio_field[i]);
            if (isNaN(numbera)) {
                pasaran_minbet_shio_field = 0;
            }
        }
        for (let i = 0; i < pasaran_maxbet_shio_field.length; i++) {
            numbera = parseInt(pasaran_maxbet_shio_field[i]);
            if (isNaN(numbera)) {
                pasaran_maxbet_shio_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limitglobal_shio_field.length; i++) {
            numbera = parseInt(pasaran_limitglobal_shio_field[i]);
            if (isNaN(numbera)) {
                pasaran_limitglobal_shio_field = 0;
            }
        }
        for (let i = 0; i < pasaran_limittotal_shio_field.length; i++) {
            numbera = parseInt(pasaran_limittotal_shio_field[i]);
            if (isNaN(numbera)) {
                pasaran_limittotal_shio_field = 0;
            }
        }
        for (let i = 0; i < pasaran_win_shio_field.length; i++) {
            numbera = parseInt(pasaran_win_shio_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_win_shio_field[i] != ".") {
                    pasaran_win_shio_field = 0;
                }
            }
        }
        for (let i = 0; i < pasaran_disc_shio_field.length; i++) {
            numbera = parseInt(pasaran_disc_shio_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_disc_shio_field[i] != ".") {
                    pasaran_disc_shio_field = 0;
                }
            }
        }
    };
    const handleKeyboard_time = (e) => {
        let numbera;
        for (let i = 0; i < pasaran_tutup_field.length; i++) {
            numbera = parseInt(pasaran_tutup_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_tutup_field[i] != ":") {
                    pasaran_tutup_field = "";
                }
            }
        }
        for (let i = 0; i < pasaran_jadwal_field.length; i++) {
            numbera = parseInt(pasaran_jadwal_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_jadwal_field[i] != ":") {
                    pasaran_jadwal_field = "";
                }
            }
        }
        for (let i = 0; i < pasaran_buka_field.length; i++) {
            numbera = parseInt(pasaran_buka_field[i]);
            if (isNaN(numbera)) {
                if (pasaran_buka_field[i] != ":") {
                    pasaran_buka_field = "";
                }
            }
        }
    };
    const removeOnline = (e) => {
        if(pasaran_tipepasaran_field != "WAJIB"){
            const pasaranonline = {
                e,
                idcomppasaran,
            };
            dispatch("handleDeletePasaranOnline", pasaranonline);
        }else{
            alert("Cannot Delete")
        }
    };
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<Container fluid style="margin-top: 70px;">
    <Row>
        <Col>
            <button
                on:click={() => {
                    BackHalaman();
                }}
                class="btn btn-dark btn-sm"
                style="border-radius: 0px;">
                Back
            </button>
        </Col>
    </Row>
    <div class="clearfix" />
    <Row>
        <Col xs="4">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Pasaran / {sData}
                    <div class="float-end">
                        {#if pasaran_tipepasaran_field != "WAJIB"}
                        <button
                            on:click={() => {
                                SaveTransaksi();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;">
                            Save
                        </button>
                        {/if}
                    </div>
                </div>
                <div class="card-body" style="height:550px;">
                    <Row>
                        <Col xs="6">
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Pasaran</label
                                >
                                <input
                                    bind:value={pasaran_name_field}
                                    type="text"
                                    disabled
                                    class="form-control"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Situs</label
                                >
                                <input
                                    bind:value={pasaran_situs_field}
                                    type="text"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Hari diundi</label
                                >
                                <input
                                    bind:value={pasaran_diundi_field}
                                    type="text"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <table>
                                    <tr>
                                        <td style="font-size: 11px;">Create</td>
                                        <td style="font-size: 11px;">:</td>
                                        <td style="font-size: 11px;"
                                            >{pasaran_create_field} - {pasaran_createdate_field}</td
                                        >
                                    </tr>
                                    <tr>
                                        <td style="font-size: 11px;">Update</td>
                                        <td style="font-size: 11px;">:</td>
                                        <td style="font-size: 11px;"
                                            >{pasaran_update_field} - {pasaran_updatedate_field}</td
                                        >
                                    </tr>
                                </table>
                            </div>
                        </Col>
                        <Col xs="6">
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >Tutup</label
                                >
                                <input
                                    bind:value={pasaran_tutup_field}
                                    on:keyup={handleKeyboard_time}
                                    minlength="0"
                                    maxlength="8"
                                    type="text"
                                    style="text-align:center;"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Jadwal</label
                                >
                                <input
                                    bind:value={pasaran_jadwal_field}
                                    on:keyup={handleKeyboard_time}
                                    minlength="0"
                                    maxlength="8"
                                    type="text"
                                    style="text-align:center;"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >Buka</label
                                >
                                <input
                                    bind:value={pasaran_buka_field}
                                    on:keyup={handleKeyboard_time}
                                    minlength="0"
                                    maxlength="8"
                                    type="text"
                                    style="text-align:center;"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Display</label
                                >
                                <input
                                    bind:value={pasaran_display_field}
                                    on:keyup={handleKeyboard_number}
                                    minlength="0"
                                    maxlength="3"
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Pasaran Togel"
                                    aria-label="Pasaran Togel"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_display_field
                                    )}</small
                                >
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">Status</label
                                >
                                <select
                                    bind:value={pasaran_status_field}
                                    class="form-control required"
                                >
                                    <option value="Y">ACTIVE</option>
                                    <option value="N">DEACTIVE</option>
                                </select>
                            </div>
                        </Col>
                    </Row>
                </div>
            </div>
        </Col>
        <Col xs="4">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Setting Pasaran Online
                    <div class="float-end">
                        {#if pasaran_tipepasaran_field != "WAJIB"}
                        <button
                            on:click={() => {
                                newOnline();
                            }}
                            class="btn btn-primary btn-sm"
                            style="border-radius: 0px;">
                            New
                        </button>
                        {/if}
                    </div>
                </div>

                <div class="card-body" style="height:550px;">
                    <table class="table table-striped">
                        <tbody>
                            {#each listPasaranOnlineByEdit as rec}
                                <tr>
                                    <td
                                        on:click={() => {
                                            removeOnline(rec.idpasaranonline);
                                        }}
                                        width="1%"
                                        NOWRAP
                                        style="text-align: center;vertical-align:top;font-size: 14px;cursor:pointer;"
                                    >
                                        <Icon name="trash" />
                                    </td>
                                    <td
                                        width="*"
                                        style="text-align: left;vertical-align:top;"
                                        >{rec.haripasaran}</td
                                    >
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        </Col>
        <Col xs="4">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Setting Limit Line + Bolak Balik
                    <div class="float-end">
                        {#if pasaran_tipepasaran_field != "WAJIB"}
                        <button
                            on:click={() => {
                                SaveLimitline();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;">
                            Save
                        </button>
                        {/if}
                    </div>
                </div>

                <div class="card-body" style="height:550px;">
                    <Row>
                        <Col xs="4">
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >BBFS</label
                                >
                                <input
                                    bind:value={pasaran_bbfs_field}
                                    type="text"
                                    disabled
                                    style="text-align:right;"
                                    class="form-control"
                                    placeholder="BBFS"
                                    aria-label="BBFS"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_bbfs_field
                                    )}</small
                                >
                            </div>
                        </Col>
                        <Col xs="4">
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >LimitLine 4D</label
                                >
                                <input
                                    bind:value={pasaran_limitline4d_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 4D"
                                    aria-label="Limitline 4D"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline4d_field
                                    )}</small
                                >
                            </div>
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >LimitLine 3D</label
                                >
                                <input
                                    bind:value={pasaran_limitline3d_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 3D"
                                    aria-label="Limitline 3D"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline3d_field
                                    )}</small
                                >
                            </div>
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >LimitLine 3DD</label
                                >
                                <input
                                    bind:value={pasaran_limitline3dd_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 3DD"
                                    aria-label="Limitline 3DD"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline3dd_field
                                    )}</small
                                >
                            </div>
                        </Col>
                        <Col xs="4">
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >LimitLine 2D</label
                                >
                                <input
                                    bind:value={pasaran_limitline2d_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 2D"
                                    aria-label="Limitline 2D"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline2d_field
                                    )}</small
                                >
                            </div>
                            <div class="mb-3">
                                <label for="exampleForm" class="form-label"
                                    >LimitLine 2DD</label
                                >
                                <input
                                    bind:value={pasaran_limitline2dd_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 2DD"
                                    aria-label="Limitline 2DD"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline2dd_field
                                    )}</small
                                >
                            </div>
                            <div class="mb-3">
                                <label
                                    for="exampleFormControlInput1"
                                    class="form-label">LimitLine 2DT</label
                                >
                                <input
                                    bind:value={pasaran_limitline2dt_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    style="text-align:right;"
                                    class="form-control required"
                                    placeholder="Limitline 2DT"
                                    aria-label="Limitline 2DT"
                                />
                                <small style="float:right;font-size:11px;"
                                    >{new Intl.NumberFormat().format(
                                        pasaran_limitline2dt_field
                                    )}</small
                                >
                            </div>
                        </Col>
                    </Row>
                </div>
            </div>
        </Col>
    </Row>
    <div class="clearfix" />
    <br />
    <Row>
        <Col>
            <TabContent pills>
                <TabPane tabId="4-3-2" tab="4D/3D/2D" active>
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        save432d();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;">
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Max Bet 4D</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet4d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 4D"
                                            aria-label="Max Bet 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet4d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Max Bet 3D</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 3D"
                                            aria-label="Max Bet 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet3d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Max Bet 3DD</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet3dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 3DD"
                                            aria-label="Max Bet 3DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet3dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Max Bet 2D</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet2d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 2D"
                                            aria-label="Max Bet 2D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet2d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Max Bet 2DD</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet2dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 2DD"
                                            aria-label="Max Bet 2DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet2dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Max Bet 2DT</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet2dt_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet 2DT"
                                            aria-label="Max Bet 2DT"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet2dt_432d_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 4D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal4d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 4D"
                                            aria-label="Limit Total 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal4d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 3D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal3d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 3DD</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal3dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3DD"
                                            aria-label="Limit Total 3DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal3dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 2D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal2d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 2D"
                                            aria-label="Limit Total 2D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal2d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 2DD</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal2dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 2DD"
                                            aria-label="Limit Total 2DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal2dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Total 2DT</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal2dt_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 2DT"
                                            aria-label="Limit Total 2DT"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal2dt_432d_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 4D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal4d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal4d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 3D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 3D"
                                            aria-label="Limit Global 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal3d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 3DD</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 3DD"
                                            aria-label="Limit Global 3DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal3dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 2D</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal2d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 2D"
                                            aria-label="Limit Global 2D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal2d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 2DD</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal2dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 2DD"
                                            aria-label="Limit Global 2DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal2dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Limit Global 2DT</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal2dt_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 2DT"
                                            aria-label="Limit Global 2DT"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal2dt_432d_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 4D(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win4d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 4D"
                                            aria-label="Win 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win4d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 3D(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 3D"
                                            aria-label="Win 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win3d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 3DD(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win3dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 3DD"
                                            aria-label="Win 3DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win3dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 2D(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win2d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 2D"
                                            aria-label="Win 2D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win2d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 2DD(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win2dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 2DD"
                                            aria-label="Win 2DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win2dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Win 2DT(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win2dt_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 2DT"
                                            aria-label="Win 2DT"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win2dt_432d_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Disc 4D(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc4d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 4D"
                                            aria-label="Disc 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc4d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Disc 3D(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc3d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 3D"
                                            aria-label="Disc 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc3d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Disc 3DD(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc3dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 3DD"
                                            aria-label="Disc 3DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc3dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label">Disc 2D(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc2d_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 2D"
                                            aria-label="Disc 2D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc2d_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Disc 2DD(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc2dd_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 2DD"
                                            aria-label="Disc 2DD"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc2dd_432d_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleFormControlInput1"
                                            class="form-label"
                                            >Disc 2DT(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc2dt_432d_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc 2DT"
                                            aria-label="Disc 2DT"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc2dt_432d_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="cbebas" tab="Colok Bebas">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveColokBebas();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Win(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Disc(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc_cbebas_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 4D"
                                            aria-label="Win 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc_cbebas_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="cmacau" tab="Colok Macau">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveColokMacau();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limitotal_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitotal_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win 2 Digit(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win2_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win2_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win 3 Digit(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win3_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win3_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win 4 Digit(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win4_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win4_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Disc(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc_cmacau_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 4D"
                                            aria-label="Win 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc_cmacau_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="cnaga" tab="Colok Naga">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveColokNaga();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win 3 Digit(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win3_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win3_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win 4 Digit(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win4_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win4_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Disc(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc_cnaga_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 4D"
                                            aria-label="Win 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc_cnaga_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="cjitu" tab="Colok Jitu">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveColokJitu();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Win AS(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_winas_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_winas_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Win Kop(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_winkop_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_winkop_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win Kepala(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_winkepala_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_winkepala_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Win Ekor(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_winekor_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Global 4D"
                                            aria-label="Limit Global 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_winekor_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Disc(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_desc_cjitu_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win 4D"
                                            aria-label="Win 4D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_desc_cjitu_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="5050umum" tab="5050 Umum">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        save5050Umum();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keibesar_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keibesar_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikecil_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikecil_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keigenap_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keigenap_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiganjil_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiganjil_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Tengah(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keitengah_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keitengah_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Tepi(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keitepi_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keitepi_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discbesar_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discbesar_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckecil_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckecil_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discgenap_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discgenap_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discganjil_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discganjil_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Tengah(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disctengah_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disctengah_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Tepi(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disctepi_5050umum_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disctepi_5050umum_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="5050special" tab="5050 Special">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        save5050Special();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_5050special_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_5050special_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei As Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiasganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei As Ganjil"
                                            aria-label="Kei As Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiasganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei As Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiasgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei As Genap"
                                            aria-label="Kei As Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiasgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei As Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiasbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei As Besar"
                                            aria-label="Kei As Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiasbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei As Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiaskecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei As Kecil"
                                            aria-label="Kei As Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiaskecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kop Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikopganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kop Ganjil"
                                            aria-label="Kei Kop Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikopganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kop Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikopgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kop Genap"
                                            aria-label="Kei Kop Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikopgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kop Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikopbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kop Besar"
                                            aria-label="Kei Kop Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikopbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kop Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikopkecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kop Kecil"
                                            aria-label="Kei Kop Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikopkecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kepala Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikepalaganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kepala Ganjil"
                                            aria-label="Kei Kepala Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikepalaganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kepala Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikepalagenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kepala Genap"
                                            aria-label="Kei Kepala Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikepalagenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kepala Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikepalabesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kepala Besar"
                                            aria-label="Kei Kepala Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikepalabesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kepala Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikepalakecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kepala Kecil"
                                            aria-label="Kei Kepala Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikepalakecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ekor Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiekorganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Ekor Ganjil"
                                            aria-label="Kei Ekor Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiekorganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ekor Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiekorgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Ekor Genap"
                                            aria-label="Kei Ekor Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiekorgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ekor Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiekorbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Ekor Besar"
                                            aria-label="Kei Ekor Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiekorbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ekor Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiekorkecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Ekor Kecil"
                                            aria-label="Kei Ekor Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiekorkecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc AS Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discasganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc AS Ganjil"
                                            aria-label="Disc AS Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discasganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc AS Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discasgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc AS Genap"
                                            aria-label="Disc AS Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discasgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc AS Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discasbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc AS Besar"
                                            aria-label="Disc AS Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discasbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc AS Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discaskecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discaskecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kop Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discasganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kop Ganjil"
                                            aria-label="Disc Kop Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discasganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kop Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckopgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kop Genap"
                                            aria-label="Disc Kop Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckopgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kop Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckopbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kop Besar"
                                            aria-label="Disc Kop Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckopbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kop Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckopkecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kop Kecil"
                                            aria-label="Kei Kop Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckopkecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kepala Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckepalaganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kepala Ganjil"
                                            aria-label="Disc Kepala Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckepalaganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kepala Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckepalagenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kepala Genap"
                                            aria-label="Disc Kepala Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckepalagenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kepala Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckepalabesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kepala Besar"
                                            aria-label="Disc Kepala Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckepalabesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kepala Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckepalakecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kepala Kecil"
                                            aria-label="Disc Kepala Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckepalakecil_5050special_field
                                            )}</small
                                        >
                                    </div>

                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ekor Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discekorganjil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Ekor Ganjil"
                                            aria-label="Disc Ekor Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discekorganjil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ekor Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discekorgenap_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Ekor Genap"
                                            aria-label="Disc Ekor Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discekorgenap_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ekor Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discekorbesar_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Ekor Besar"
                                            aria-label="Disc Ekor Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discekorbesar_5050special_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ekor Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discekorkecil_5050special_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Ekor Kecil"
                                            aria-label="Disc Ekor Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discekorkecil_5050special_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="5050kombinasi" tab="5050 Kombinasi">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        save5050Kombinasi();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Kei Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangkeimono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Kei Mono"
                                            aria-label="Belakang Kei Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangkeimono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Kei Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangkeistereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Kei Stereo"
                                            aria-label="Belakang Kei Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangkeistereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Kei Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangkeikembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Kei Kembang"
                                            aria-label="Belakang Kei Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangkeikembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Kei Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangkeikempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Kei Kempis"
                                            aria-label="Belakang Kei Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangkeikempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Kei Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangkeikembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Kei Kembar"
                                            aria-label="Belakang Kei Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangkeikembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Kei Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahkeimono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Kei Mono"
                                            aria-label="Tengah Kei Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahkeimono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Kei Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahkeistereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Kei Stereo"
                                            aria-label="Tengah Kei Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahkeistereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Kei Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahkeikembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Kei Kembang"
                                            aria-label="Tengah Kei Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahkeikembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Kei Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahkeikempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Kei Kempis"
                                            aria-label="Tengah Kei Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahkeikempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Kei Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahkeikembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Kei Kembar"
                                            aria-label="Tengah Kei Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahkeikembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Kei Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depankeimono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Kei Mono"
                                            aria-label="Depan Kei Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depankeimono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Kei Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depankeistereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Kei Stereo"
                                            aria-label="Depan Kei Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depankeistereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Kei Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depankeikembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Kei Kembang"
                                            aria-label="Depan Kei Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depankeikembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Kei Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depankeikempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Kei Kempis"
                                            aria-label="Depan Kei Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depankeikempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Kei Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depankeikembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Kei Kembar"
                                            aria-label="Depan Kei Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depankeikembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Disc Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangdiscmono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Disc Mono"
                                            aria-label="Belakang Disc Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangdiscmono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Disc Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangdiscstereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Disc Stereo"
                                            aria-label="Belakang Disc Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangdiscstereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Disc Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangdisckembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Disc Kembang"
                                            aria-label="Belakang Disc Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangdisckembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Disc Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangdisckempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Disc Kempis"
                                            aria-label="Belakang Disc Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangdisckempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Belakang Disc Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_belakangdisckembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Belakang Disc Kembar"
                                            aria-label="Belakang Disc Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_belakangdisckembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Disc Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahdiscmono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Disc Mono"
                                            aria-label="Tengah Disc Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahdiscmono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Disc Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahdiscstereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Disc Stereo"
                                            aria-label="Tengah Disc Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahdiscstereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Disc Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahdisckembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Disc Kembang"
                                            aria-label="Tengah Disc Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahdisckembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Disc Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahdisckempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Disc Kempis"
                                            aria-label="Tengah Disc Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahdisckempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Tengah Disc Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_tengahdisckembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Tengah Disc Kembar"
                                            aria-label="Tengah Disc Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_tengahdisckembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Disc Mono(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depandiscmono_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Disc Mono"
                                            aria-label="Depan Disc Mono"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depandiscmono_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Disc Stereo(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depandiscstereo_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Disc Stereo"
                                            aria-label="Depan Disc Stereo"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depandiscstereo_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>

                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Disc Kembang(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depandisckembang_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Disc Kembang"
                                            aria-label="Depan Disc Kembang"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depandisckembang_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Disc Kempis(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depandisckempis_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Disc Kempis"
                                            aria-label="Depan Disc Kempis"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depandisckempis_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Depan Disc Kembar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_depandisckembar_5050kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Depan Disc Kembar"
                                            aria-label="Depan Disc Kembar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_depandisckembar_5050kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="macaukombinasi" tab="Macau/Kombinasi">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveMacauKombinasi();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Win(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win"
                                            aria-label="Win"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Discount(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc_kombinasi_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Discount"
                                            aria-label="Discount"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc_kombinasi_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="dasar" tab="Dasar">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveDasar();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_dasar_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_dasar_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keibesar_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Besar"
                                            aria-label="Kei Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keibesar_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keikecil_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Kecil"
                                            aria-label="Kei Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keikecil_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keigenap_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Genap"
                                            aria-label="Kei Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keigenap_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Kei Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_keiganjil_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Kei Ganjil"
                                            aria-label="Kei Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_keiganjil_dasar_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Besar(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discbesar_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Besar"
                                            aria-label="Disc Besar"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discbesar_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Kecil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disckecil_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Kecil"
                                            aria-label="Disc Kecil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disckecil_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Genap(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discgenap_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Genap"
                                            aria-label="Disc Genap"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discgenap_dasar_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Disc Ganjil(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_discganjil_dasar_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc Ganjil"
                                            aria-label="Disc Ganjil"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_discganjil_dasar_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
                <TabPane tabId="shio" tab="Shio">
                    <Panel height_body="700px">
                        <slot:template slot="cheader">
                            <div class="float-end">
                                {#if pasaran_tipepasaran_field != "WAJIB"}
                                <button
                                    on:click={() => {
                                        saveShio();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;"
                                >
                                    Save
                                </button>
                                {/if}
                            </div>
                        </slot:template>
                        <slot:template slot="cbody">
                            <Row>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <select
                                            bind:value={pasaran_shioyear_shio_field}
                                            name="selectshio"
                                            class="form-control required"
                                        >
                                            <option value="kambing"
                                                >Kambing</option
                                            >
                                            <option value="kuda">Kuda</option>
                                            <option value="ular">Ular</option>
                                            <option value="naga">Naga</option>
                                            <option value="kelinci"
                                                >Kelinci</option
                                            >
                                            <option value="harimau"
                                                >Harimau</option
                                            >
                                            <option value="kerbau"
                                                >Kerbau</option
                                            >
                                            <option value="tikus">Tikus</option>
                                            <option value="babi">Babi</option>
                                            <option value="anjing"
                                                >Anjing</option
                                            >
                                            <option value="ayam">Ayam</option>
                                            <option value="monyet"
                                                >Monyet</option
                                            >
                                        </select>
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Min Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_minbet_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Min Bet"
                                            aria-label="Min Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_minbet_shio_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Max Bet</label
                                        >
                                        <input
                                            bind:value={pasaran_maxbet_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Max Bet"
                                            aria-label="Max Bet"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_maxbet_shio_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Total</label
                                        >
                                        <input
                                            bind:value={pasaran_limittotal_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total"
                                            aria-label="Limit Total"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limittotal_shio_field
                                            )}</small
                                        >
                                    </div>
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label"
                                            >Limit Global</label
                                        >
                                        <input
                                            bind:value={pasaran_limitglobal_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Limit Total 3D"
                                            aria-label="Limit Total 3D"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_limitglobal_shio_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Win(x)</label
                                        >
                                        <input
                                            bind:value={pasaran_win_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Win"
                                            aria-label="Win"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_win_shio_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2">
                                    <div class="mb-3">
                                        <label
                                            for="exampleForm"
                                            class="form-label">Disc(%)</label
                                        >
                                        <input
                                            bind:value={pasaran_disc_shio_field}
                                            on:keyup={handleKeyboard_number}
                                            type="text"
                                            style="text-align:right;"
                                            class="form-control required"
                                            placeholder="Disc"
                                            aria-label="Disc"
                                        />
                                        <small
                                            style="float:right;font-size:11px;"
                                            >{new Intl.NumberFormat().format(
                                                pasaran_disc_shio_field
                                            )}</small
                                        >
                                    </div>
                                </Col>
                                <Col xs="2" />
                            </Row>
                        </slot:template>
                    </Panel>
                </TabPane>
            </TabContent>
        </Col>
    </Row>
</Container>
<!-- Modal -->
<div
    class="modal fade"
    id="modaNewOnline"
    tabindex="-1"
    aria-labelledby="exampleModalLabel"
    aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">
                    Setting Pasaran Online
                </h5>
                <button
                    type="button"
                    class="btn-close"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                />
            </div>
            <div class="modal-body">
                <select bind:value={select_pasaranonline} class="form-control">
                    <option value="senin">Senin</option>
                    <option value="selasa">Selasa</option>
                    <option value="rabu">Rabu</option>
                    <option value="kamis">Kamis</option>
                    <option value="jumat">Jumat</option>
                    <option value="sabtu">Sabtu</option>
                    <option value="minggu">Minggu</option>
                </select>
            </div>
            <div class="modal-footer">
                <button
                    on:click={() => {
                        savePasaranOnline();
                    }}
                    type="button"
                    class="btn btn-warning btn-sm">SAVE</button>
            </div>
        </div>
    </div>
</div>
