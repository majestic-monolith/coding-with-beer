<script lang="ts">
    import { type ChartConfiguration } from "chart.js";
  import { Chart } from "chart.js/auto";
  
  let chartObject: Chart;

  const config: ChartConfiguration = {
    type: "line",
    data: {
      labels: [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
      ],
      datasets: [
        {
          label: new Date().getFullYear().toString(),
          backgroundColor: "#4c51bf",
          borderColor: "#4c51bf",
          data: [65, 78, 66, 44, 56, 67, 75],
          fill: false
        },
        {
          label: (new Date().getFullYear() - 1).toString(),
          fill: false,
          backgroundColor: "#fff",
          borderColor: "#fff",
          data: [40, 68, 86, 74, 56, 60, 87]
        }
      ]
    },
    options: {
      maintainAspectRatio: false,
      responsive: true,
      plugins: {
        title: {
          display: false,
          text: "Sales Charts",
          color: "white",
        },
        legend: {
          labels: {
            color: "white",
          },
          align: "end",
          position: "bottom",
        },
        tooltip: {
          mode: "index",
          intersect: false,
        },
      },
      scales: {
        x: {
          ticks: {
            color: "rgba(255,255,255,.7)",
          },
          display: true,
          title: {
            display: false,
            text: "Month",
            color: "white",
          },
          grid: {
            display: false,
            color: "rgba(33, 37, 41, 0.3)",
            tickBorderDash: [2],
            tickBorderDashOffset: 2,
          },
        },
        y: {
          ticks: {
            color: "rgba(255,255,255,.7)",
          },
          display: true,
          title: {
            display: false,
            text: "Value",
            color: "white",
          },
          grid: {
            color: "rgba(255, 255, 255, 0.15)",
            tickBorderDash: [2],
            tickBorderDashOffset: 2,
          },
        },
      },
    },
  };

  function chart(node: HTMLCanvasElement) {
    function setupChart() {
      chartObject = new Chart(node, config);
    }
		
    setupChart();
  }
</script>
<div
  class="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded bg-blueGray-700"
>
  <div class="rounded-t mb-0 px-4 py-3 bg-transparent">
    <div class="flex flex-wrap items-center">
      <div class="relative w-full max-w-full flex-grow flex-1">
        <h6 class="uppercase text-blueGray-100 mb-1 text-xs font-semibold">
          Overview
        </h6>
        <h2 class="text-white text-xl font-semibold">
          Sales value
        </h2>
      </div>
    </div>
  </div>
  <div class="p-4 flex-auto">
    <!-- Chart -->
    <div class="relative h-350-px">
      <canvas use:chart></canvas>
    </div>
  </div>
</div>
