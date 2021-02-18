import "./Sidebar.css";

const Sidebar = ({ sidebarOpen, closeSidebar }) => {
    return(
        <div className={sidebarOpen ? "sidebar-responsive" : ""} id="sidebar">
            
            <div className="sidbar__title">
                <div className="sidebar_img">
                    <img src={} alt="logo" />
                    <h1>Contact Manager</h1>
                </div>
                <i className="fa fa-times" id="sidebarIcon" onClick={() => closeSidebar()}></i>
            </div>

            <div className="sidebar__menu">
                <div className="sidebar__link active_menu_link">
                    <i className="fa fa-home"></i>
                    <a href="#">Dashboard</a>
                </div> 
            </div>

        </div>
    );
}

export default Sidebar;