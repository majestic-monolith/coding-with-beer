<script>
    // components for this layout
    import AdminNavbar from "$lib/components/Navbars/AdminNavbar.svelte";
    import Sidebar from "$lib/components/Sidebar/Sidebar.svelte";
    import HeaderStats from "$lib/components/Headers/HeaderStats.svelte";
    import FooterAdmin from "$lib/components/Footers/FooterAdmin.svelte";
    
    let { children } = $props();
  </script>
  
  <div>
    <Sidebar/>
    <div class="relative md:ml-64 bg-blueGray-100">
      <AdminNavbar />
      <HeaderStats />
      <div class="px-4 md:px-10 mx-auto w-full -m-24">
        {@render children()}
        <FooterAdmin />
      </div>
    </div>
  </div>
  