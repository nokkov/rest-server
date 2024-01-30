import './App.css';

function App() {
  return (
    <div className="App">
      <div className="container">
        <div className="logo">
          <img src="ULTIMATE-URL-SHORTENER.gif" alt="logo" />
        </div>
      </div>
      <div className="input-fields">
        <form id="urlForm">
          <div className="input-fields-wrapper">
            <input type="text" id="longUrlField" required></input>
            <button type="button" onClick="submitForm()">
              SEND NUDES! 
            </button>
          </div>
        </form>  
      </div>
      <script defer src="./handler.js"></script>
    </div>
  );
}

export default App;
