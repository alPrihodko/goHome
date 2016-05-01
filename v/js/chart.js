
var chart;
var data;
var options;

google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart);

      function drawChart() {
        data = new google.visualization.DataTable();

        data.addColumn('datetime', 'XAxis');
        data.addColumn('number', 'Internal temp');

        options = {
          title: 'Temperature chart',
          curveType: 'function',
          legend: { position: 'bottom' }
        };

        chart = new google.visualization.LineChart(document.getElementById('curve_chart'));

        fetch('/control/hdata')
          .then(function(response) {
            if (response.status == 200) {
              // Examine the text in the response

              response.json().then(function(dt) {
                console.log(dt);
                var ar = new Array();
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
          })
          /*
          .then(function(resp) {
            console.log(resp);
            var o = JSON.parse(resp);
            for (var i=0; i < o.length; i++) {
              var r = o[i];
              data.addRows(r.Timestamp*1000, r.TempInside);
            }
            chart.draw(data, options);
          })
          */

        //chart.draw(data, options);
      };

      function updateChart(t) {
            data.addRows([[new Date(), t]]);
            chart.draw(data, options);
      }
