<script>
    import dayjs from "dayjs";
    import isBetween from "dayjs/plugin/isBetween";
    import { Row, Col, Container, Input } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    dayjs.extend(isBetween);

    let token = localStorage.getItem("token");
    let akses_page = true;
    let listwinlose = [];
    let total_turnover = 0;
    let total_winlose = 0;
    let total_winlose_agent = 0;
    let css_sub_winlosemember = "color:red;font-weight:bold;";
    let css_sub_winloseagent = "color:red;font-weight:bold;";
    let start_field = "";
    let end_field = "";
    let css_loader = "display: none;";
    let msgloader = "";
    async function initapp() {
        const res = await fetch("/api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "REPORTWINLOSE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        }
    }
    async function call_generatereport() {
        listwinlose = [];
        total_turnover = 0;
        total_winlose = 0;
        total_winlose_agent = 0;
        let flag = false;
        let date1 = dayjs(start_field);
        let date2 = dayjs(end_field);
        const diff = date2.diff(date1, "day", true);
        const days = Math.floor(diff);

        if (days > 0 && days < 31) {
            flag = true;
        }
        if (flag) {
            const res = await fetch("/api/reportwinlose", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    client_start: start_field,
                    client_end: end_field,
                }),
            });
            const json = await res.json();
            let record = json.record;
            if (json.status == 200) {
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        let css_winlosemember = "color:red;font-weight:bold;";
                        let css_winloseagent = "color:red;font-weight:bold;";
                        if (record[i]["report_client_winlose"] > 0) {
                            css_winlosemember = "color:blue;font-weight:bold;";
                        }
                        if (record[i]["report_agent_winlose"] > 0) {
                            css_winloseagent = "color:blue;font-weight:bold;";
                        }
                        listwinlose = [
                            ...listwinlose,
                            {
                                report_client_username:
                                    record[i]["report_client_username"],
                                report_client_turnover:
                                    record[i]["report_client_turnover"],
                                report_client_winlose:
                                    record[i]["report_client_winlose"],
                                report_client_winlose_css: css_winlosemember,
                                report_agent_winlose:
                                    record[i]["report_agent_winlose"],
                                report_agent_winlose_css: css_winloseagent,
                            },
                        ];
                    }
                    total_turnover = parseInt(json.subtotalturnover);
                    total_winlose = parseInt(json.subtotalwinlose);
                    total_winlose_agent = parseInt(json.subtotalwinlosecompany);

                    if (total_winlose > 0) {
                        css_sub_winlosemember = "color:blue;font-weight:bold;";
                    }
                    if (total_winlose_agent > 0) {
                        css_sub_winloseagent = "color:blue;font-weight:bold;";
                    }
                }
            }
        } else {
            alert(
                "Ups Generate, hany bisa dilakukan dengan range tanggal 31 hari"
            );
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
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
                        Report Winlose
                        <div class="float-end">
                            <button
                                on:click={() => {
                                    call_generatereport();
                                }}
                                class="btn btn-warning"
                                style="border-radius: 0px;"
                            >
                                Generate
                            </button>
                        </div>
                    </div>
                    <div class="card-body" style="height:450px;">
                        <Row>
                            <div class="mb-3">
                                <label for="example" class="form-label"
                                    >Start</label
                                >
                                <Input
                                    bind:value={start_field}
                                    type="date"
                                    name="date"
                                    id="exampleDate"
                                    data-date-format="dd-mm-yyyy"
                                    placeholder="date placeholder"
                                />
                            </div>
                            <div class="mb-3">
                                <label for="example" class="form-label"
                                    >End</label
                                >
                                <Input
                                    bind:value={end_field}
                                    type="date"
                                    name="date"
                                    id="exampleDate"
                                    data-date-format="dd-mm-yyyy"
                                    placeholder="date placeholder"
                                />
                            </div>
                        </Row>
                    </div>
                </div>
            </Col>
            <Col xs="9">
                <Panel
                    height_body="500px"
                    css_footer="padding:10px;margin:0px;"
                >
                    <slot:template slot="cheader">
                        Report Winlose
                    </slot:template>
                    <slot:template slot="cbody">
                        <table class="table">
                            <thead>
                                <tr>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align: top;font-size: 13px;"
                                        >USERNAME</th
                                    >

                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >TURNOVER</th
                                    >
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >WINLOSE MEMBER</th
                                    >
                                    <th
                                        width="20%"
                                        style="text-align: right;vertical-align: top;font-size: 13px;"
                                        >WINLOSE AGENT</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each listwinlose as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align: top;font-size: 12px;"
                                            >{rec.report_client_username}</td
                                        >

                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;color:blue;"
                                            >{new Intl.NumberFormat().format(
                                                rec.report_client_turnover
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;{rec.report_client_winlose_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.report_client_winlose
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align: top;font-size: 12px;{rec.report_agent_winlose_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.report_agent_winlose
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
                                        >TOTAL TURNOVER</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;font-weight: bold;color:blue;font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            total_turnover
                                        )}</td
                                    >
                                </tr>
                                <tr style="padding: 0px;margin:0px;">
                                    <td
                                        style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none"
                                        >MEMBER WINLOSE</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;{css_sub_winlosemember}font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            total_winlose
                                        )}</td
                                    >
                                </tr>
                                <tr style="padding: 0px;margin:0px;">
                                    <td
                                        style="text-align: left;vertical-align:top;font-weight: bold;font-size:12px;border:none"
                                        >AGENT WINLOSE</td
                                    >
                                    <td
                                        style="text-align: right;vertical-align:top;{css_sub_winloseagent}font-size:12px;border:none;"
                                        >{new Intl.NumberFormat().format(
                                            total_winlose_agent
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
