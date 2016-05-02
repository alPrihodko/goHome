
var chart;
var data;
var options;
var DATALIMIT = 5000;
var ChartTimeLimit = 60*60;

function setChartTimeLimit(t) {
  if (t === "1 Hour") {
    window.ChartTimeLimit = 60*24*60;
  } else if (t === "1 Day") {
    window.ChartTimeLimit = 60*24*30*60;
  } else {
    window.ChartTimeLimit = 60*60;
  }
}

google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart);

      function drawChart() {
        data = new google.visualization.DataTable();

        data.addColumn('datetime', 'XAxis');
        data.addColumn('number', 'Internal temp');

        options = {
          title: 'Temperature chart',
          curveType: 'function',
          legend: { position: 'bottom' },
          page: 'enable',
          pageSize: 20
        };

        chart = new google.visualization.LineChart(document.getElementById('curve_chart'));
        loadChart();
      };

      function updateChart(t) {
            data.addRow([new Date(), t]);
            chart.draw(data, options);
            if (data.getNumberOfRows() > DATALIMIT) {
              data.removeRows(0, 5);
            }
      }

      function loadChart() {
        data.removeRows(0, data.getNumberOfRows());
        fetch('/control/hdata?from=' + ChartTimeLimit)
          .then(function(response) {
            if (response.status == 200) {
              // Examine the text in the response

              response.json().then(function(dt) {
                //console.log(dt);
                var ar = [];
                for (var i=0; i < dt.length; i++) {
                  var r = dt[i];
                  ar.push([new Date(r.Timestamp*1000), parseFloat(r.TempInside)]);
                }
                data.addRows(ar);
                chart.draw(data, options);
              });

            } else {
              alert(response.statusText)
            }
          });
      }
