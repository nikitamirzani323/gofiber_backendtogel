<script>
    import Home from "../periode/Home.svelte";
    import Entry from "../periode/Entry.svelte";

    let listPeriode = [];
    let listPeriodePasaran = [];
    let record = "";
    let recordpasaran = "";
    let totalrecord = 0;
    let idtrxkeluaran = "";
    let idpasarancode = "";
    let sData = "";
    let akses_page = true;
    let token = localStorage.getItem("token");
    let periode_tglkeluaran_field = "";
    let periode_tanggalnext_field = "";
    let periode_periode_field = "";
    let periode_keluaran_field = "";
    let periode_status_field = "";
    let periode_statusonline_field = "";
    let periode_statusrevisi_field = "";
    let periode_create_field = "";
    let periode_createdate_field = "";
    let periode_update_field = "";
    let periode_updatedate_field = "";

    const handleRefreshData = (e) => {
        listPeriode = [];
        listPeriodePasaran = [];
        totalrecord = 0;
        setTimeout(function () {
            initPeriode();
        }, 1000);
    };
    const handleBackHalaman = () => {
        idtrxkeluaran = "";
        idpasarancode = "";
        sData = "";
        periode_tglkeluaran_field = "";
        periode_tanggalnext_field = "";
        periode_periode_field = "";
        periode_keluaran_field = "";
        periode_status_field = "";
        periode_create_field = "";
        periode_createdate_field = "";
        periode_update_field = "";
        periode_updatedate_field = "";
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        idtrxkeluaran = e.detail.e;
        idpasarancode = e.detail.f;
        sData = "Edit";

        editPeriode(idtrxkeluaran);
    };
    const handleRefreshEdit = (e) => {
        idtrxkeluaran = e.detail;
        sData = "Edit";
        editPeriode(idtrxkeluaran);
    };
    async function initapp() {
        const res = await fetch("/api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "PERIODE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status == 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            // setTimeout(function(){ initPasaran() }, 1000);
            initPeriode();
        }
    }
    async function initPeriode() {
        const res = await fetch("/api/allperiode", {
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
            recordpasaran = json.pasaranonline;
            let css_totalmember = "";
            let css_totalbet = "";
            let css_totalbayar = "";
            let css_winlose = "";
            if (record != null) {
                totalrecord = record.length;
                for (var i = 0; i < record.length; i++) {
                    if (parseInt(record[i]["pasaran_totalmember"]) > 0) {
                        css_totalmember = "color:blue;font-weight:bold";
                    } else {
                        css_totalmember = "color:red;font-weight:bold";
                    }
                    if (parseInt(record[i]["pasaran_totalbet"]) > 0) {
                        css_totalbet = "color:blue;font-weight:bold";
                    } else {
                        css_totalbet = "color:red;font-weight:bold";
                    }
                    if (parseInt(record[i]["pasaran_totaloutstanding"]) > 0) {
                        css_totalbayar = "color:blue;font-weight:bold";
                    } else {
                        css_totalbayar = "color:red;font-weight:bold";
                    }
                    if (parseInt(record[i]["pasaran_winlose"]) > 0) {
                        css_winlose = "color:blue;font-weight:bold";
                    } else {
                        css_winlose = "color:red;font-weight:bold";
                    }
                    listPeriode = [
                        ...listPeriode,
                        {
                            no: record[i]["pasaran_no"],
                            pasaran_invoice: record[i]["pasaran_invoice"],
                            pasaran_idcompp: record[i]["pasaran_idcompp"],
                            pasaran_code: record[i]["pasaran_code"],
                            pasaran_periode: record[i]["pasaran_periode"],
                            pasaran_name: record[i]["pasaran_name"],
                            pasaran_keluaran: record[i]["pasaran_keluaran"],
                            pasaran_tanggal: record[i]["pasaran_tanggal"],
                            pasaran_totalmember:
                                record[i]["pasaran_totalmember"],
                            pasaran_totalmember_css: css_totalmember,
                            pasaran_totalbet: record[i]["pasaran_totalbet"],
                            pasaran_totalbet_css: css_totalbet,
                            pasaran_totaloutstanding:
                                record[i]["pasaran_totaloutstanding"],
                            pasaran_totaloutstanding_css: css_totalbayar,
                            pasaran_winlose: record[i]["pasaran_winlose"],
                            pasaran_winlose_css: css_winlose,
                            pasaran_status: record[i]["pasaran_status"],
                            pasaran_status_css: record[i]["pasaran_status_css"],
                        },
                    ];
                }
            }
            if (recordpasaran != null) {
                for (var j = 0; j < recordpasaran.length; j++) {
                    listPeriodePasaran = [
                        ...listPeriodePasaran,
                        {
                            pasarancomp_idcompp:
                                recordpasaran[j]["pasarancomp_idcompp"],
                            pasarancomp_nama:
                                recordpasaran[j]["pasarancomp_nama"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function editPeriode(e) {
        const res = await fetch("/api/editperiode", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(e),
            }),
        });
        const json = await res.json();
        let record = json.record;

        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                let keluaran = record[i]["periode_keluaran"];
                let document_status = "OPEN";
                if (keluaran != "") {
                    document_status = "LOCK";
                }
                periode_status_field = document_status;
                periode_tglkeluaran_field =
                    record[i]["periode_tanggalkeluaran"];
                periode_tanggalnext_field = record[i]["periode_tanggalnext"];
                periode_periode_field = record[i]["periode_keluaranperiode"];
                periode_keluaran_field = record[i]["periode_keluaran"];
                periode_statusrevisi_field = record[i]["periode_statusrevisi"];
                periode_statusonline_field = record[i]["periode_statusonline"];
                periode_create_field = record[i]["periode_create"];
                periode_createdate_field = record[i]["periode_createdate"];
                periode_update_field = record[i]["periode_update"];
                periode_updatedate_field = record[i]["periode_updatedate"];
            }
        }
    }

    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>

{#if sData == "" && idtrxkeluaran == ""}
    {#if akses_page == true}
        <Home
            on:handleRefreshData={handleRefreshData}
            on:handleEditData={handleEditData}
            {token}
            {totalrecord}
            {listPeriode}
            {listPeriodePasaran}
        />
    {/if}
{:else}
    <Entry
        on:handleBackHalaman={handleBackHalaman}
        on:handleRefreshEdit={handleRefreshEdit}
        {sData}
        {token}
        {idtrxkeluaran}
        {idpasarancode}
        {periode_status_field}
        {periode_tglkeluaran_field}
        {periode_tanggalnext_field}
        {periode_periode_field}
        {periode_keluaran_field}
        {periode_statusonline_field}
        {periode_statusrevisi_field}
        {periode_create_field}
        {periode_createdate_field}
        {periode_update_field}
        {periode_updatedate_field}
    />
{/if}
