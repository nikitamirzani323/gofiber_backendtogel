<script>
    import { Row, Col, Container } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    let token = localStorage.getItem("token");
    let akses_page = true;
    let listPasaranCompany = [];
    let listresult = [];
    let prediksi_select_field = "";
    let prediksi_field = "";
    let css_loader = "display: none;";
    let msgloader = "";
    let totalbet = 0;
    let totalwin = 0;
    let subtotal = 0;
    let subtotal_css = "";
    async function callPrediksi() {
        listresult = [];
        totalbet = 0;
        totalwin = 0;
        subtotal = 0;
        subtotal_css = "";
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (prediksi_select_field == "") {
            flag = true;
            msg += "The Pasaran is required\n";
        }
        if (prediksi_field == "") {
            flag = true;
            msg += "The Nomor Keluaran is required\n";
        }
        if (parseInt(prediksi_field.length) < 4) {
            flag = true;
            msg += "The Nomor Keluaran is must 4 Character\n";
        }
        if (flag == false) {
            const res = await fetch("/api/listprediksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    nomorkeluaran: prediksi_field,
                    pasaran_code: parseInt(prediksi_select_field),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
                let record = json.record;
                totalbet = json.totalbet;
                totalwin = json.subtotal;
                subtotal = json.subtotalwin;
                if (parseInt(subtotal) > 0) {
                    subtotal_css = "color:blue;";
                } else {
                    subtotal_css = "color:red;";
                }
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listresult = [
                            ...listresult,
                            {
                                prediksi_invoice: record[i]["prediksi_invoice"],
                                prediksi_code: record[i]["prediksi_code"],
                                prediksi_tanggal: record[i]["prediksi_tanggal"],
                                prediksi_username: record[i]["prediksi_username"],
                                prediksi_permainan: record[i]["prediksi_permainan"],
                                prediksi_posisitogel: record[i]["prediksi_posisitogel"],
                                prediksi_nomor: record[i]["prediksi_nomor"],
                                prediksi_bet: record[i]["prediksi_bet"],
                                prediksi_diskon: record[i]["prediksi_diskon"],
                                prediksi_diskonpercen: record[i]["prediksi_diskonpercen"].toFixed(2),
                                prediksi_kei: record[i]["prediksi_kei"],
                                prediksi_keipercen: record[i]["prediksi_keipercen"].toFixed(2),
                                prediksi_bayar: record[i]["prediksi_bayar"],
                                prediksi_win: record[i]["prediksi_win"].toFixed(2),
                                prediksi_totalwin: record[i]["prediksi_totalwin"],
                                prediksi_status: record[i]["prediksi_status"],
                                prediksi_statuscss: record[i]["prediksi_statuscss"],
                            },
                        ];
                    }
                }
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }
    }
    async function initapp() {
        const res = await fetch("/api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "PREDIKSI-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            initprediksi();
        }
    }
    async function initprediksi() {
        const res = await fetch("/api/listpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({}),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listPasaranCompany = [
                        ...listPasaranCompany,
                        {
                            pasaran_idcomp: record[i]["pasaran_idcomp"],
                            pasaran_code: record[i]["pasaran_code"],
                            pasaran_name: record[i]["pasaran_name"],
                        },
                    ];
                }
            } else {
                alert("Error");
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleKeyboard_number = (e) => {
        let numbera;
		for (let i = 0; i < prediksi_field.length; i++) {
			numbera = parseInt(prediksi_field[i]);
			if (isNaN(numbera)) {
				prediksi_field = "";
			}
		}
    }
    initapp();
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
{#if akses_page == true}
    <Container fluid style="margin-top: 70px;">
        <Row>
            <Col xs="3">
                <div class="card" style="border-radius: 0px;margin-top:10px;">
                    <div class="card-header" style="">
                        Prediksi
                        <div class="float-end">
                            <button
                                on:click={() => {
                                    callPrediksi();
                                }}
                                class="btn btn-warning btn-sm"
                                style="border-radius: 0px;">
                                Check
                            </button>
                        </div>
                    </div>
                    <div class="card-body" style="height:450px;">
                        <Row>
                            <div class="mb-3">
                                <label for="example" class="form-label">Pasaran</label>
                                <select
                                    bind:value={prediksi_select_field}
                                    placeholder="Pilih Pasaran"
                                    class="form-control">
                                    {#each listPasaranCompany as rec}
                                        <option value={rec.pasaran_idcomp}
                                            >{rec.pasaran_name}</option
                                        >
                                    {/each}
                                </select>
                            </div>
                            <div class="mb-3">
                                <label for="example" class="form-label">Nomor Keluaran</label>
                                <input
                                    bind:value={prediksi_field}
                                    on:keyup={handleKeyboard_number}
                                    type="text"
                                    minlength="4"
                                    maxlength="4"
                                    class="form-control"
                                    placeholder="Nomor Keluaran: Ex: 1234"
                                    aria-label="Nomor Keluaran"/>
                            </div>
                        </Row>
                    </div>
                </div>
            </Col>
            <Col xs="9">
                <Panel
                    height_body="700px"
                    css_footer="padding:10px;margin:0px;">
                    <slot:template slot="cheader"> List Bet </slot:template>
                    <slot:template slot="cbody">
                        <table class="table">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;"
                                        >STATUS</th
                                    >
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">INVOICE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">CODE</th>
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align: top;font-size: 13px;">TANGGAL</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;"
                                        >USERNAME</th
                                    >
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">TIPE</th>
                                    <th
                                        width="7%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;">PERMAINAN</th>
                                    <th
                                        width="1%"
                                        style="text-align: left;vertical-align: top;font-size: 13px;"
                                        >NOMOR</th
                                    >
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >BET</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >DISC</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >KEI</th
                                    >
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >BAYAR</th
                                    >
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >WIN</th
                                    >
                                    <th
                                        width="7%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >WIN<br />TOTAL</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each listresult as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align: top;font-size: 12px;{rec.prediksi_statuscss}"
                                            >{rec.prediksi_status}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.prediksi_invoice}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.prediksi_code}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align: top;font-size: 12px;">{rec.prediksi_tanggal}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;"
                                            >{rec.prediksi_username}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.prediksi_posisitogel}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;">{rec.prediksi_permainan}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;"
                                            >{rec.prediksi_nomor}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;"
                                            >{new Intl.NumberFormat().format(
                                                rec.prediksi_bet
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;">
                                            {rec.prediksi_diskon}&nbsp;({rec.prediksi_diskonpercen}%)
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;">
                                            {rec.prediksi_kei}&nbsp;({rec.prediksi_keipercen}%)
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                            >{new Intl.NumberFormat().format(
                                                rec.prediksi_bayar
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;"
                                            >{rec.prediksi_win}x</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;color:red;"
                                            >{new Intl.NumberFormat().format(
                                                rec.prediksi_totalwin
                                            )}</td
                                        >
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </slot:template>
                    <slot:template slot="cfooter">
                        <table
                            width="100%"
                            style="padding:0px;margin-bottom:0px;"
                        >
                            <tbody>
                                <tr style="padding: 0px;margin:0px;">
                                    <td
                                        style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none;"
                                        >TOTAL BET</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            totalbet
                                        )}</td
                                    >
                                </tr>
                                <tr style="padding: 0px;margin:0px;">
                                    <td
                                        style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none"
                                        >MEMBER WINLOSE</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            totalwin
                                        )}</td
                                    >
                                </tr>
                                <tr style="padding: 0px;margin:0px;">
                                    <td
                                        style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none"
                                        >COMPANY WINLOSE</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;font-weight: bold;{subtotal_css};font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            subtotal
                                        )}</td
                                    >
                                </tr>
                            </tbody>
                        </table>
                    </slot:template>
                </Panel>
            </Col>
        </Row>
    </Container>
{/if}
