package gerador

const htmlTemplate = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Relatório Executivo</title>
    <style>
        :root {
            --bg-main: #f8f9fa;
            --text-main: #212529;
            --primary: #0d6efd;
            --table-header: #212529;
            --table-border: #dee2e6;
        }
        body {
            font-family: system-ui, -apple-system, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
            background-color: var(--bg-main);
            color: var(--text-main);
            margin: 0;
            padding: 2rem;
            line-height: 1.5;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
        }
        header {
            margin-bottom: 2rem;
            border-bottom: 2px solid var(--primary);
            padding-bottom: 1rem;
        }
        h1 { margin: 0; color: var(--table-header); font-size: 2rem; }
        .timestamp { color: #6c757d; font-size: 0.9rem; margin-top: 0.5rem; }
        
        .card {
            background: #fff;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.05);
            padding: 1.5rem;
            margin-bottom: 2rem;
        }
        h2 { margin-top: 0; color: var(--primary); font-size: 1.4rem; border-bottom: 1px solid var(--table-border); padding-bottom: 0.5rem;}
        
        .table-responsive {
            overflow-x: auto;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 1rem;
            font-size: 0.95rem;
        }
        th, td {
            padding: 0.75rem 1rem;
            text-align: left;
            border-bottom: 1px solid var(--table-border);
        }
        th {
            background-color: var(--table-header);
            color: #fff;
            font-weight: 600;
            text-transform: uppercase;
            font-size: 0.8rem;
            letter-spacing: 0.5px;
        }
        tr:hover { background-color: #f1f3f5; }
        .no-data { color: #6c757d; font-style: italic; padding: 1rem; text-align: center; }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>Painel de Relatórios</h1>
            <div class="timestamp">Gerado em: {{.GeneratedAt}}</div>
        </header>

        {{range $reportName, $table := .Reports}}
        <div class="card">
            <h2>{{$reportName}}</h2>
            <div class="table-responsive">
                {{if $table.Rows}}
                <table>
                    <thead>
                        <tr>
                            {{range $table.Headers}}
                            <th>{{.}}</th>
                            {{end}}
                        </tr>
                    </thead>
                    <tbody>
                        {{range $row := $table.Rows}}
                        <tr>
                            {{range $key := $table.Headers}}
                            <td>{{index $row $key}}</td>
                            {{end}}
                        </tr>
                        {{end}}
                    </tbody>
                </table>
                {{else}}
                <div class="no-data">Nenhum dado processado por este pipeline.</div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>
`
