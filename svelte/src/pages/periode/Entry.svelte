<script>
    import { Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Panel from "../../components/Panel.svelte";
    import Modal from "../../components/Modal.svelte";

    export let sData = "";
    export let idtrxkeluaran = 0;
    export let idpasarancode = "";
    export let token = "";
    export let periode_status_field = "LOCK";
    export let periode_tglkeluaran_field = "";
    export let periode_tanggalnext_field = "";
    export let periode_periode_field = "";
    export let periode_keluaran_field = "";
    export let periode_statusonline_field = "";
    export let periode_statusrevisi_field = "";
    export let periode_create_field = "";
    export let periode_createdate_field = "";
    export let periode_update_field = "";
    export let periode_updatedate_field = "";
    let listBetTableGroup = [];
    let listBetTable = [];
    let listBet = [];
    let listBetStatus = [];
    let listBetUsername = [];
    let listMember = [];
    let listMemberNomor = [];
    let client_username = "";
    let totalbet = 0;
    let totalbayar = 0;
    let totalwin = 0;
    let subtotal_member_bet = 0;
    let subtotal_member_bayar = 0;
    let subtotal_member_cancel = 0;
    let subtotal_member_win = 0;
    let subtotal_member_winlose = 0;
    let css_winlose = "color:red;font-weight:bold;";
    let chooce_permainan = "";
    let css_loader = "display: none;";
    let msgloader = "";
    let myModalrevisi = "";
    let dispatch = createEventDispatcher();

    

    const schema = yup.object().shape({
        msgrevisi: yup
            .string()
            .required()
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Alasan revisi must Character A-Z or a-z or 1-9 "
            )
            .min(5, "Alasan revisi must be at least 5 Character")
            .max(120, "Alasan revisi must be at least 120 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            msgrevisi: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            revisiTransaksi(values.msgrevisi);
        },
    });

    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };
    async function SaveTransaksi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (periode_keluaran_field == "") {
            flag = true;
            msg += "The keluaran is required\n";
        }
        if (parseInt(periode_keluaran_field.length) < 4) {
            flag = true;
            msg += "The keluaran is must 4 Character\n";
        }

        if (flag == false) {
            periode_status_field = "LOCK";
            const res = await fetch("/api/saveperiode", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sData: sData,
                    page: "PERIODE-SAVE",
                    idinvoice: parseInt(idtrxkeluaran),
                    idpasarancode: idpasarancode,
                    nomorkeluaran: periode_keluaran_field,
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
                periode_keluaran_field = "";
            } else {
                msgloader = json.message;
            }
            dispatch("handleRefreshEdit", idtrxkeluaran);
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
            listBet = [];
            listMember = [];
            call_listbet("4D");
            call_listmember();
        } else {
            alert(msg);
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }
    }
    function callrevisiTransaksi() {
        myModalrevisi = new bootstrap.Modal(
            document.getElementById("modalrevisi")
        );
        myModalrevisi.show();
    }
    async function revisiTransaksi(e) {
        periode_status_field = "LOCK";
        const res = await fetch("/api/saveperioderevisi", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sData: sData,
                page: "PERIODE-SAVE",
                idinvoice: parseInt(idtrxkeluaran),
                msgrevisi: e,
            }),
        });
        const json = await res.json();

        if (json.status == 200) {
            msgloader = json.message;
        } else if (json.status == 403) {
            alert(json.message);
            periode_keluaran_field = "";
        } else {
            msgloader = json.message;
        }
        $form.msgrevisi = "";
        dispatch("handleRefreshEdit", idtrxkeluaran);
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        listBet = [];
        listMember = [];
        call_listbet("4D");
        call_listmember();
        myModalrevisi.hide();
    }
    async function cancelbetTransaksi(e,f) {
        const res = await fetch("/api/cancelbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sData: sData,
                page: "PERIODE-SAVE",
                permainan: f,
                idinvoice: parseInt(idtrxkeluaran),
                idinvoicedetail: parseInt(e),
            }),
        });
        const json = await res.json();

        if (json.status == 200) {
            msgloader = json.message;
        } else if (json.status == 403) {
            alert(json.message);
            periode_keluaran_field = "";
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
        listBet = [];
        listMember = [];
        call_listbet("4D");
        call_listmember();
    }
    async function call_listmembernomor(nomor) {
        const res = await fetch("/api/periodelistmemberbynomor", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: chooce_permainan,
                nomor: nomor,
            }),
        });
        const json = await res.json();
        let record = json.record;
        let nomember = 0;

        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                let myModal = new bootstrap.Modal(
                    document.getElementById("modalmemberbet")
                );
                myModal.show();
                for (var i = 0; i < record.length; i++) {
                    nomember = nomember + 1;
                    listMemberNomor = [
                        ...listMemberNomor,
                        {
                            member_no: nomember,
                            member_name: record[i]["member"],
                            member_permainan: record[i]["member_permainan"],
                            member_nomor: record[i]["member_nomor"],
                            member_bet: record[i]["member_bet"],
                            member_disc: record[i]["member_disc"],
                            member_discpercen: record[i]["member_discpercen"],
                            member_kei: record[i]["member_kei"],
                            member_keipercen: record[i]["member_keipercen"],
                            member_bayar: record[i]["member_bayar"],
                            member_win: record[i]["member_win"],
                            member_winhasil: record[i]["member_winhasil"],
                        },
                    ];
                }
            } else {
                setTimeout(function () {
                    msgloader = "";
                    css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listmember() {
        const res = await fetch("/api/periodelistmember", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
            }),
        });
        const json = await res.json();
        let record = json.record;
        let nomember = 0;

        if (json.status === 400) {
            logout();
        } else {
            subtotal_member_bet = 0;
            subtotal_member_bayar = 0;
            subtotal_member_cancel = 0;
            subtotal_member_win = 0;
            subtotal_member_winlose = 0;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    nomember = nomember + 1;
                    subtotal_member_bet =
                        subtotal_member_bet + parseInt(record[i]["totalbet"]);
                    subtotal_member_bayar =
                        subtotal_member_bayar +
                        parseInt(record[i]["totalbayar"]);
                    subtotal_member_cancel =
                        subtotal_member_cancel +
                        parseInt(record[i]["totalcancelbet"]);
                    subtotal_member_win =
                        subtotal_member_win + parseInt(record[i]["totalwin"]);
                    listMember = [
                        ...listMember,
                        {
                            member_no: nomember,
                            member_name: record[i]["member"],
                            member_totalbet: record[i]["totalbet"],
                            member_totalbayar: record[i]["totalbayar"],
                            member_totalcancel: record[i]["totalcancelbet"],
                            member_totalwin: record[i]["totalwin"],
                        },
                    ];
                }
                subtotal_member_winlose =
                    subtotal_member_bayar -
                    subtotal_member_cancel -
                    subtotal_member_win;
                if (parseInt(subtotal_member_winlose) > 0) {
                    css_winlose = "color:blue;font-weight:bold;";
                }
            } else {
                setTimeout(function () {
                    msgloader = "";
                    css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listbet(e) {
        listBet = [];
        const res = await fetch("/api/periodelistbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["bet_status"] != "CANCEL") {
                        listBet = [
                            ...listBet,
                            {
                                bet_id: record[i]["bet_id"].toString(),
                                bet_datetime: record[i]["bet_datetime"],
                                bet_ipaddress: record[i]["bet_ipaddress"],
                                bet_device: record[i]["bet_device"],
                                bet_timezone: record[i]["bet_timezone"],
                                bet_username: record[i]["bet_username"],
                                bet_typegame: record[i]["bet_typegame"],
                                bet_nomortogel: record[i]["bet_nomortogel"],
                                bet_bet: record[i]["bet_bet"],
                                bet_diskon: record[i]["bet_diskon"],
                                bet_diskonpercen: record[i]["bet_diskonpercen"],
                                bet_kei: record[i]["bet_kei"],
                                bet_keipercen: record[i]["bet_keipercen"],
                                bet_bayar: record[i]["bet_bayar"],
                                bet_win: record[i]["bet_win"],
                                bet_totalwin: record[i]["bet_totalwin"],
                                bet_status: record[i]["bet_status"],
                                bet_statuscss: record[i]["bet_statuscss"],
                            },
                        ];
                    }
                }
                // call_bettable();
            } else {
                setTimeout(function () {
                    msgloader = "";
                    css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listgroupbet(e) {
        listBetStatus = [];
        const res = await fetch("/api/periodelistbetstatus", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                status: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetStatus = [
                        ...listBetStatus,
                        {
                            bet_id: record[i]["bet_id"],
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: record[i]["bet_statuscss"],
                        },
                    ];
                }
            } else {
                setTimeout(function () {
                    msgloader = "";
                    css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listbetbyusername(e) {
        listBetUsername = [];
        client_username = e.toUpperCase();
        const res = await fetch("/api/periodelistbetusername", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                username: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetUsername = [
                        ...listBetUsername,
                        {
                            bet_id: record[i]["bet_id"],
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"],
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"],
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"],
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_statuscss: record[i]["bet_statuscss"],
                        },
                    ];
                }
                let myModal = new bootstrap.Modal(
                    document.getElementById("modallistbetusername")
                );
                myModal.show();
            } else {
                setTimeout(function () {
                    msgloader = "";
                    css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listbettable() {
        const res = await fetch("/api/periodelistbettable", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetTable = [
                        ...listBetTable,
                        {
                            permainan: record[i]["permainan"],
                        },
                    ];
                }
            }
        }
    }
    async function call_bettable(e) {
        chooce_permainan = e;
        const res = await fetch("/api/periodebettable", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
            logout();
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetTableGroup = [
                        ...listBetTableGroup,
                        {
                            bet_keluaran: record[i]["bet_keluaran"],
                            bet_totalbet: record[i]["bet_totalbet"],
                            bet_totalmember: record[i]["bet_totalmember"],
                        },
                    ];
                }
            }
        }
    }
    const handleKeyboard_number = (e) => {
        let numbera;
        for (let i = 0; i < periode_keluaran_field.length; i++) {
            numbera = parseInt(periode_keluaran_field[i]);
            if (isNaN(numbera)) {
                periode_keluaran_field = "";
            }
        }
    };
    let searchBet = "";
    let filteritems = [];
    let searchBetUsername = "";
    let filteritemsusername = [];

    $: {
        if (searchBet) {
            filteritems = listBet.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchBet.toLowerCase()) ||
                    item.bet_nomortogel
                        .toLowerCase()
                        .includes(searchBet.toLowerCase()) ||
                    item.bet_typegame
                        .toLowerCase()
                        .includes(searchBet.toLowerCase()) ||
                    item.bet_id.toLowerCase().includes(searchBet.toLowerCase())
            );
        } else {
            filteritems = [...listBet];
        }
        if (searchBetUsername) {
            filteritemsusername = listBetUsername.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchBetUsername.toLowerCase()) ||
                    item.bet_nomortogel
                        .toLowerCase()
                        .includes(searchBetUsername.toLowerCase()) ||
                    item.bet_typegame
                        .toLowerCase()
                        .includes(searchBetUsername.toLowerCase())
            );
        } else {
            filteritemsusername = [...listBetUsername];
        }
        if ($errors.msgrevisi) {
            alert($errors.msgrevisi);
            $form.msgrevisi = "";
        }
    }
    const cancelBet = (e,f) => {
        cancelbetTransaksi(e,f);
    };
    const listbetbyusername = (e) => {
        call_listbetbyusername(e);
    };
    const openCity = (e) => {
        listBetTableGroup = [];
        call_bettable(e);
    };
    const groupMember = (nomor) => {
        listMemberNomor = [];
        call_listmembernomor(nomor);
    };
    const handleSelectPermainan = (event) => {
        if (event.target.value != "") {
            call_listbet(event.target.value);
        } else {
            listBet = [];
        }
    };

    call_listmember();
    call_listbettable();
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
        <Col xs="3">
            <div
                class="alert alert-warning"
                role="alert"
                style="margin-top: 10px;">
                <h4>INFORMATION</h4>
                <p>Periode Selanjutnya : {periode_tanggalnext_field}</p>
            </div>
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Periode / {sData}
                    <div class="float-end">
                        {#if periode_status_field == "OPEN"}
                            {#if periode_statusonline_field == "OFFLINE"}
                                <button
                                    on:click={() => {
                                        SaveTransaksi();
                                    }}
                                    class="btn btn-warning btn-sm"
                                    style="border-radius: 0px;">
                                    Save
                                </button>
                            {/if}
                        {:else if periode_statusrevisi_field == "OPEN"}
                            <button
                                on:click={() => {
                                    callrevisiTransaksi();
                                }}
                                class="btn btn-warning btn-sm"
                                style="border-radius: 0px;">
                                Revisi
                            </button>
                        {/if}
                    </div>
                </div>
                <div class="card-body" style="height:450px;">
                    <Row>
                        <div class="mb-3">
                            <label
                                for="exampleFormControlInput1"
                                class="form-label">Invoice</label
                            >
                            <input
                                bind:value={idtrxkeluaran}
                                type="text"
                                disabled
                                class="form-control"
                                placeholder="Invoice"
                                aria-label="Invoice"
                            />
                        </div>
                        <div class="mb-3">
                            <label
                                for="exampleFormControlInput1"
                                class="form-label">Tanggal</label
                            >
                            <input
                                bind:value={periode_tglkeluaran_field}
                                type="text"
                                style="text-align: center;"
                                disabled
                                class="form-control"
                                placeholder="Tanggal"
                                aria-label="Tanggal"
                            />
                        </div>
                        <div class="mb-3">
                            <label
                                for="exampleFormControlInput1"
                                class="form-label">Periode</label
                            >
                            <input
                                bind:value={periode_periode_field}
                                type="text"
                                disabled
                                class="form-control required"
                                placeholder="Periode"
                                aria-label="Periode"
                            />
                        </div>
                        <div class="mb-3">
                            <label
                                for="exampleFormControlInput1"
                                class="form-label">Keluaran</label
                            >
                            <input
                                bind:value={periode_keluaran_field}
                                on:keyup={handleKeyboard_number}
                                type="text"
                                maxlength="4"
                                class="form-control required"
                                placeholder="Keluaran"
                                aria-label="Keluaran"
                            />
                        </div>
                        <div class="mb-3">
                            <table>
                                <tr>
                                    <td style="font-size: 11px;">Create</td>
                                    <td style="font-size: 11px;">:</td>
                                    <td style="font-size: 11px;"
                                        >{periode_create_field} - {periode_createdate_field}</td
                                    >
                                </tr>
                                {#if periode_update_field != ""}
                                    <tr>
                                        <td style="font-size: 11px;">Update</td>
                                        <td style="font-size: 11px;">:</td>
                                        <td style="font-size: 11px;"
                                            >{periode_update_field} - {periode_updatedate_field}</td
                                        >
                                    </tr>
                                {/if}
                            </table>
                        </div>
                    </Row>
                </div>
            </div>
            <div class="clearfix" />
            <br />
        </Col>
        <Col xs="4">
            <Panel height_body="500px" css_footer="padding:10px;margin:0px;">
                <slot:template slot="cheader"> List Member </slot:template>
                <slot:template slot="cbody">
                    <table class="table">
                        <thead>
                            <tr>
                                <th
                                    width="1%"
                                    style="text-align: center;vertical-align: top;font-size: 13px;">NO</th>
                                <th
                                    width="*"
                                    style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                <th
                                    NOWRAP
                                    width="20%"
                                    style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL<br />BET</th>
                                <th
                                    NOWRAP
                                    width="20%"
                                    style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL<br />BAYAR</th>
                                <th
                                    NOWRAP
                                    width="20%"
                                    style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL<br />CANCEL</th>
                                <th
                                    NOWRAP
                                    width="20%"
                                    style="text-align: right;vertical-align: top;font-size: 13px;">TOTAL<br />WIN</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each listMember as rec}
                                <tr>
                                    <td
                                        NOWRAP
                                        style="text-align: center;vertical-align: top;font-size: 12px;">{rec.member_no}</td>
                                    <td
                                        on:click={() => {
                                            listbetbyusername(rec.member_name);
                                        }}
                                        NOWRAP
                                        style="text-decoration:underline;cursor:pointer;text-align: left;vertical-align: top;font-size: 12px;">{rec.member_name}</td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                        {new Intl.NumberFormat().format(
                                            rec.member_totalbet
                                        )}
                                    </td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                        {new Intl.NumberFormat().format(
                                            rec.member_totalbayar
                                        )}
                                    </td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                        {new Intl.NumberFormat().format(
                                            rec.member_totalcancel
                                        )}
                                    </td>
                                    <td
                                        NOWRAP
                                        style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                        {new Intl.NumberFormat().format(
                                            rec.member_totalwin
                                        )}
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </slot:template>
                <slot:template slot="cfooter">
                    <table width="100%" style="padding:0px;margin-bottom:0px;">
                        <tbody>
                            <tr style="padding: 0px;margin:0px;">
                                <td
                                    style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none;">TOTAL BET</td>
                                <td
                                    style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                    {new Intl.NumberFormat().format(
                                        subtotal_member_bet
                                    )}
                                </td>
                            </tr>
                            <tr style="padding: 0px;margin:0px;">
                                <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL BAYAR</td>
                                <td style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                    {new Intl.NumberFormat().format(
                                        subtotal_member_bayar
                                    )}
                                </td>
                            </tr>
                            <tr style="padding: 0px;margin:0px;">
                                <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL CANCEL</td>
                                <td style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                    {new Intl.NumberFormat().format(
                                        subtotal_member_cancel
                                    )}
                                </td>
                            </tr>
                            <tr style="padding: 0px;margin:0px;">
                                <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL WIN</td>
                                <td style="text-align: right;vertical-align:top;font-weight: bold;color:red;font-size:12px;border:none;">
                                    {new Intl.NumberFormat().format(
                                        subtotal_member_win
                                    )}
                                </td>
                            </tr>
                            <tr style="padding: 0px;margin:0px;">
                                <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL WINLOSE</td>
                                <td
                                    style="text-align: right;vertical-align:top;{css_winlose}font-size:12px;border:none;">
                                    {new Intl.NumberFormat().format(
                                        subtotal_member_winlose
                                    )}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </slot:template>
            </Panel>
        </Col>
        <Col xs="5">
            <Panel height_body="510px" css_footer="">
                <slot:template slot="cheader"> Bet Group </slot:template>
                <slot:template slot="cbody">
                    <div class="d-flex align-items-start">
                        <div
                            class="nav flex-column nav-pills me-3"
                            id="v-pills-tab"
                            role="tablist"
                            aria-orientation="vertical">
                            {#each listBetTable as rec}
                                <button
                                    on:click={() => {
                                        openCity(rec.permainan);
                                    }}
                                    style="font-size:13px;"
                                    class="nav-link"
                                    id="listbet_4d-tab"
                                    data-bs-toggle="pill"
                                    data-bs-target="#listbet_4d"
                                    type="button"
                                    role="tab"
                                    aria-controls="v-pills-home"
                                    aria-selected="true">{rec.permainan}</button>
                            {/each}
                        </div>
                        <div
                            class="tab-content"
                            id="v-pills-tabContent"
                            style="width:100%;">
                            <table class="table">
                                <thead>
                                    <tr>
                                        <th style="text-align: center;vertical-align: top;font-size:13px;">NOMOR</th>
                                        <th style="text-align: right;vertical-align: top;font-size:13px;">TOTAL MEMBER</th>
                                        <th style="text-align: right;vertical-align: top;font-size:13px;">TOTAL BET</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each listBetTableGroup as rec}
                                        <tr>
                                            <td style="text-align: center;vertical-align: top;font-size:12px;">{rec.bet_keluaran}</td>
                                            <td
                                                on:click={() => {
                                                    groupMember(
                                                        rec.bet_keluaran
                                                    );
                                                }}
                                                style="text-decoration:underline;cursor:pointer;text-align: right;vertical-align: top;font-size:12px;">{rec.bet_totalmember}</td>
                                            <td style="text-align: right;vertical-align: top;font-size:12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_totalbet
                                                )}
                                            </td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </slot:template>
            </Panel>
        </Col>
        <div class="clearfix" />
        <Col xs="12">
            <ul class="nav nav-pills mb-3" id="pills-tab" role="tablist">
                <li
                    on:click={() => {
                        call_listgroupbet("all");
                    }}
                    class="nav-item"
                    role="presentation">
                    <button
                        class="nav-link active"
                        id="pills-all-tab"
                        data-bs-toggle="pill"
                        data-bs-target="#pills-all"
                        type="button"
                        role="tab"
                        aria-controls="pills-all"
                        aria-selected="true">ALL</button>
                </li>
                <li
                    on:click={() => {
                        call_listgroupbet("winner");
                    }}
                    class="nav-item"
                    role="presentation">
                    <button
                        class="nav-link"
                        id="pills-winner-tab"
                        data-bs-toggle="pill"
                        data-bs-target="#pills-winner"
                        type="button"
                        role="tab"
                        aria-controls="pills-winner"
                        aria-selected="false">WINNER</button>
                </li>
                <li
                    on:click={() => {
                        call_listgroupbet("cancel");
                    }}
                    class="nav-item"
                    role="presentation">
                    <button
                        class="nav-link"
                        id="pills-cancel-tab"
                        data-bs-toggle="pill"
                        data-bs-target="#pills-cancel"
                        type="button"
                        role="tab"
                        aria-controls="pills-cancel"
                        aria-selected="false">CANCEL</button>
                </li>
            </ul>
            <div class="tab-content" id="pills-tabContent">
                <div
                    class="tab-pane fade show active"
                    id="pills-all"
                    role="tabpanel"
                    aria-labelledby="pills-all-tab">
                    <Panel
                        height_body="700px"
                        css_footer="padding:10px;margin:0px;">
                        <slot:template slot="cheader">List Bet</slot:template>
                        <slot:template slot="cbody">
                            <div class="row">
                                <div class="col-lg-5">
                                    <select
                                        on:change={handleSelectPermainan}
                                        class="form-control">
                                        <option value="">--Pilih Permainan--</option>
                                        {#each listBetTable as rec}
                                            <option value={rec.permainan}>{rec.permainan}</option>
                                        {/each}
                                    </select>
                                </div>
                                <div class="col-lg-7">
                                    <input
                                        class="form-control"
                                        placeholder="Searching"
                                        bind:value={searchBet}
                                        type="text"/>
                                </div>
                            </div>
                            <table class="table" width="100%">
                                <thead>
                                    <tr>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">&nbsp;</th>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                        <th
                                            width="10%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                        <th
                                            width="*"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                        <th
                                            width="1%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each filteritems as rec}
                                        <tr>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 11px;">
                                                {#if periode_keluaran_field == ""}
                                                    {#if rec.bet_status == "RUNNING"}
                                                        <button
                                                            on:click={() => {
                                                                cancelBet(
                                                                    rec.bet_id,rec.bet_typegame
                                                                );
                                                            }}
                                                            class="btn btn-danger btn-sm"
                                                            ><i
                                                                class="bi bi-trash"
                                                            /></button
                                                        >
                                                    {/if}
                                                {/if}
                                            </td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}">{rec.bet_status}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_id}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_datetime}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_username}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_ipaddress}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_device}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_timezone}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_typegame}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_nomortogel}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;"
                                                >{new Intl.NumberFormat().format(
                                                    rec.bet_bet
                                                )}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                                >{new Intl.NumberFormat().format(
                                                    rec.bet_bayar
                                                )}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;">{rec.bet_win}x</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                                >{new Intl.NumberFormat().format(
                                                    rec.bet_totalwin
                                                )}</td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </slot:template>
                        <slot:template slot="cfooter">
                            <table
                                width="100%"
                                style="padding:0px;margin-bottom:0px;">
                                <tbody>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none;">TOTAL BET</td>
                                        <td style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbet
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL BAYAR</td>
                                        <td style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbayar
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL WIN</td>
                                        <td style="text-align: right;vertical-align:top;font-weight: bold;color:red;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalwin
                                            )}
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </slot:template>
                    </Panel>
                </div>
                <div
                    class="tab-pane fade"
                    id="pills-winner"
                    role="tabpanel"
                    aria-labelledby="pills-winner-tab">
                    <Panel
                        height_body="700px"
                        css_footer="padding:10px;margin:0px;">
                        <slot:template slot="cheader"> List Bet </slot:template>
                        <slot:template slot="cbody">
                            <table class="table" width="100%">
                                <thead>
                                    <tr>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                        <th
                                            width="10%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                        <th
                                            width="*"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                        <th
                                            width="1%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each listBetStatus as rec}
                                        <tr>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}">{rec.bet_status}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_id}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_datetime}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_username}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_ipaddress}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_device}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_timezone}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_typegame}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_nomortogel}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_bet
                                                )}
                                            </td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_bayar
                                                )}
                                            </td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;">{rec.bet_win}x</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_totalwin
                                                )}
                                            </td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </slot:template>
                        <slot:template slot="cfooter">
                            <table
                                width="100%"
                                style="padding:0px;margin-bottom:0px;">
                                <tbody>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none;">TOTAL BET</td>
                                        <td
                                            style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbet
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL BAYAR</td>
                                        <td
                                            style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbayar
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL WIN</td>
                                        <td
                                            style="text-align: right;vertical-align:top;font-weight: bold;color:red;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalwin
                                            )}
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </slot:template>
                    </Panel>
                </div>
                <div
                    class="tab-pane fade"
                    id="pills-cancel"
                    role="tabpanel"
                    aria-labelledby="pills-cancel-tab">
                    <Panel
                        height_body="700px"
                        css_footer="padding:10px;margin:0px;">
                        <slot:template slot="cheader"> List Bet </slot:template>
                        <slot:template slot="cbody">
                            <table class="table" width="100%">
                                <thead>
                                    <tr>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                                        <th
                                            width="1%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                                        <th
                                            width="10%"
                                            style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                        <th
                                            width="*"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                                        <th
                                            width="7%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                        <th
                                            width="1%"
                                            style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                                        <th
                                            width="10%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                                        <th
                                            width="20%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                                        <th
                                            width="7%"
                                            style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each listBetStatus as rec}
                                        <tr>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}">{rec.bet_status}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_id}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_datetime}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_username}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_ipaddress}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_device}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_timezone}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_typegame}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_nomortogel}</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_bet
                                                )}
                                            </td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_bayar
                                                )}
                                            </td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;">{rec.bet_win}x</td>
                                            <td
                                                NOWRAP
                                                style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                                {new Intl.NumberFormat().format(
                                                    rec.bet_totalwin
                                                )}
                                            </td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </slot:template>
                        <slot:template slot="cfooter">
                            <table
                                width="100%"
                                style="padding:0px;margin-bottom:0px;">
                                <tbody>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none;">TOTAL BET</td>
                                        <td
                                            style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbet
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL BAYAR</td>
                                        <td
                                            style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalbayar
                                            )}
                                        </td>
                                    </tr>
                                    <tr style="padding: 0px;margin:0px;">
                                        <td style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none">TOTAL WIN</td>
                                        <td style="text-align: right;vertical-align:top;font-weight: bold;color:red;font-size:12px;border:none;">
                                            {new Intl.NumberFormat().format(
                                                totalwin
                                            )}
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </slot:template>
                    </Panel>
                </div>
            </div>
        </Col>
    </Row>
</Container>

<Modal
    modal_id={"modalmemberbet"}
    modal_size={"modal-xl modal-dialog-centered"}
    modal_body_height={"height:500px;overflow:scroll;"}
    modal_footer_flag={false}>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">INFORMATION</h5>
    </slot:template>
    <slot:template slot="body">
        <table class="table">
            <thead>
                <tr>
                    <th
                        width="*"
                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                    <th
                        width="1%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                </tr>
            </thead>
            <tbody>
                {#each listMemberNomor as rec}
                    <tr>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.member_name}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.member_permainan}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.member_nomor}</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;">
                            {new Intl.NumberFormat().format(
                                rec.member_bet
                            )}
                        </td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{rec.member_disc}&nbsp;({rec.member_discpercen}%)</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{rec.member_kei}&nbsp;({rec.member_keipercen}%)</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                            {new Intl.NumberFormat().format(
                                rec.member_bayar
                            )}
                        </td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;">{rec.member_win}x</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                            {new Intl.NumberFormat().format(
                                rec.member_winhasil
                            )}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>
<Modal
    modal_id={"modalrevisi"}
    modal_size={"modal-dialog-centered"}
    modal_body_height={"height:150px;"}
    modal_footer_flag={true}>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModal">REVISI</h5>
    </slot:template>
    <slot:template slot="body">
        <Container>
            <Row>
                <div class="mb-3">
                    <input
                        on:change={handleChange}
                        bind:value={$form.msgrevisi}
                        invalid={$errors.msgrevisi.length > 0}
                        type="text"
                        maxlength="120"
                        class="form-control required"
                        placeholder="Alasan Revisi"
                        aria-label="Alasan Revisi"
                    />
                </div>
            </Row>
        </Container>
    </slot:template>
    <slot:template slot="footer">
        <div class="float-end">
            <button
                on:click={() => {
                    handleSubmit();
                }}
                class="btn btn-warning"
                style="border-radius: 0px;">
                Save
            </button>
        </div>
    </slot:template>
</Modal>

<Modal
    modal_id={"modallistbetusername"}
    modal_size={"modal-xl modal-dialog-centered"}
    modal_body_height={"height:500px;overflow:scroll;"}
    modal_footer_flag={false}>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">
            INFORMATION : {client_username}
        </h5>
    </slot:template>
    <slot:template slot="body">
        <div class="col-lg-12" style="padding: 5px;">
            <input
                class="form-control"
                placeholder="Searching"
                bind:value={searchBetUsername}
                type="text"/>
        </div>
        <table class="table" width="100%">
            <thead>
                <tr>
                    <th
                        width="1%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">STATUS</th>
                    <th
                        width="1%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">CODE</th>
                    <th
                        width="10%"
                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                    <th
                        width="*"
                        style="text-align: left;vertical-align: top;font-size: 13px;">USERNAME</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">IPADDRESS</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">BROWSER</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">TIMEZONE</th>
                    <th
                        width="7%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                    <th
                        width="1%"
                        style="text-align: left;vertical-align: top;font-size: 13px;">NOMOR</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BET</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">DISC</th>
                    <th
                        width="10%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">KEI</th>
                    <th
                        width="20%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">BAYAR</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN</th>
                    <th
                        width="7%"
                        style="text-align: right;vertical-align: top;font-size: 13px;">WIN<br />TOTAL</th>
                </tr>
            </thead>
            <tbody>
                {#each filteritemsusername as rec}
                    <tr>
                        <td
                            NOWRAP
                            style="text-align: center;vertical-align: top;font-size: 12px;{rec.bet_statuscss}">{rec.bet_status}</td>
                        <td
                            NOWRAP
                            style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_id}</td>
                        <td
                            NOWRAP
                            style="text-align: center;vertical-align: top;font-size: 12px;">{rec.bet_datetime}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_username}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_ipaddress}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_device}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_timezone}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_typegame}</td>
                        <td
                            NOWRAP
                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.bet_nomortogel}</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;">{rec.bet_diskon}&nbsp;({rec.bet_diskonpercen}%)</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">{rec.bet_kei}&nbsp;({rec.bet_keipercen}%)</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                            {new Intl.NumberFormat().format(rec.bet_bayar)}
                        </td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;">{rec.bet_win}x</td>
                        <td
                            NOWRAP
                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                            {new Intl.NumberFormat().format(
                                rec.bet_totalwin
                            )}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>
