<script lang="ts">
    // components for this layout
    import AuthNavbar from "$lib/components/Navbars/AuthNavbar.svelte";
    import FooterSmall from "$lib/components/Footers/FooterSmall.svelte";
  
    let { children } = $props();
    const registerBg2 = "/img/register_bg_2.png";
  </script>
  
  <div>
    <AuthNavbar />
    <main>
      <section class="relative w-full h-full py-40 min-h-screen">
        <div
          class="absolute top-0 w-full h-full bg-blueGray-800 bg-no-repeat bg-full"
          style="background-image: url({registerBg2});"
        ></div> 
        {@render children()}
        <FooterSmall absolute={true} />
      </section>
    </main>
  </div>
  