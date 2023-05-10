
import React from 'react';
import UserList from '../../components/user';
import Navbar from '../../components/navbar';

function Users() {
    return ( 
      <div className="App">
      <Navbar/>
      <div style={{marginTop: "50px"}}>
      <UserList/>
      </div>
      
    </div>
    
     );
}

export default Users;