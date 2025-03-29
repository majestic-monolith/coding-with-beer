<script lang="ts">
    // core components
    import CardLineChart from "$lib/components/Cards/CardLineChart.svelte";
    import CardBarChart from "$lib/components/Cards/CardBarChart.svelte";
    import CardPageVisits from "$lib/components/Cards/CardPageVisits.svelte";
    import CardSocialTraffic from "$lib/components/Cards/CardSocialTraffic.svelte";
</script>

<div>
    <div class="flex flex-wrap">
      <div class="w-full xl:w-8/12 mb-12 xl:mb-0 px-4">
        <CardLineChart />
      </div>
      <div class="w-full xl:w-4/12 px-4">
        <CardBarChart />
      </div>
    </div>
    <div class="flex flex-wrap mt-4">
      <div class="w-full xl:w-8/12 mb-12 xl:mb-0 px-4">
        <CardPageVisits />
      </div>
      <div class="w-full xl:w-4/12 px-4">
        <CardSocialTraffic />
      </div>
    </div>
  </div>
  