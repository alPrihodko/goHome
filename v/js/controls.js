
var Ns = function(){};

Ns.ws = new WebSocket("ws://192.168.1.20:1234/relays");
Ns.DataHandler = new Object();
Ns.Heat = new Object();
Ns.Pump = new Object();


Ns.DataHandler.handle = function (msg) {
  //alert(msg)
  var o = JSON.parse(msg);
  switch (o.Type) {
    case "pumpStateChanged":
    Ns.Pump.callback(o);
    break;

    case "heatStateChanged":
    Ns.Heat.callback(o);
    break;

  }
};

Ns.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    Ns.DataHandler.handle(event.data);
};


var ButtonMotor = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/pump')
      .then(function(response) {
        //alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
        //alert(response.status); // 200
        if (response.status == 200) {
          //self.setState({message: "Pump: " + response.text});
          var txt = response.text();
          //alert (txt)
          return txt;
        } else {
          alert(response.statusText)
        }
       })
       .then(function(state) {
         self.props.val = state;
         self.setState({message: "Pump: " + state});
         //alert(state);
       })
    return {message: "checking..."};
  },

  handleSubmit: function() {
    /*
      N.ws.send(this.props.val);
      N.ws.send("{\"Command\":\"" + N.CommandOnPumpr1 + "\"}");
      */
      if (this.props.val === "Auto") {
        this.props.val = "On";
      } else if (this.props.val === "On") {
        this.props.val = "Off";
      } else {
        this.props.val = "Auto";
      }
      var v = this.props.val
      var self = this
      fetch('/control/pump?state=' + v)
        .then(function(response) {
          //alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
          //alert(response.status); // 200
          if (response.status == 200) {
            //self.setState({message: "Pump: " + v});
            return response.text();
          } else {
            alert(response.statusText)
          }
         })
         .then(function(state) {
           //st = JSON.parse(state);
           //if (st.Type == "pumpStateChange") {
             //self.props.val = st.Value;
             //self.setState({message: self.props.val});
          // }
           //alert(st.toString);
         })
        .catch( alert );
  },

  handleChange(e) {
     this.setState({message: e.target.value})
   },

   componentWillMount() {
     Ns.Pump.callback = (data) => {
       switch (data.Key ) {
         case "state":
           this.setState({
             message: "Pump: " + data.Value
           });
           break;
       }
     };
   },


  render: function() {
    return (
      <div>
      <button className="mui-btn mui-btn--flat bdefault" onClick={this.handleSubmit} >
      {this.state.message}</button>
      </div>
    );
  }
});


var ButtonHeater = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/heat')
      .then(function(response) {

        if (response.status == 200) {
          return response.text();
        } else {
          alert(response.statusText)
        }
       })
       .then(function(state) {
         self.props.val = state;
         self.setState({message: "Heater: " + state});
         //alert(state);
       })
    return {message: "checking..."};
  },

  handleSubmit: function() {
      if (this.props.val === "Auto") {
        this.props.val = "On";
      } else if (this.props.val === "On") {
        this.props.val = "Off";
      } else {
        this.props.val = "Auto";
      }
      this.setState({message: "Heater: " + this.props.val})

      var v = this.props.val
      var self = this
      fetch('/control/heat?state=' + v)
        .then(function(response) {

          if (response.status == 200) {
            //self.setState({message: "Heater: " + v});
            return response;
          } else {
            alert(response.statusText)
          }
         })
          .catch( alert );
  },

  handleChange(e) {
     this.setState({message: e.target.value})
   },

   componentWillMount() {
     Ns.Heat.callback = (data) => {
       switch (data.Key ) {
         case "state":
           this.setState({
             message: "Heat: " + data.Value
           });
           break;
       }
     };
   },

   /*
  componentWillMount(){
    DataHandler.callback = (data) => {
       this.setState({message: data.message});
     };
  },
*/
  render: function() {
    return (
      <div>
      <button className="mui-btn mui-btn--flat bdefault" onClick={this.handleSubmit} >
      {this.state.message}</button>
      </div>
    );
  }
});



ReactDOM.render(<ButtonMotor msg="Pump: Auto" />, document.getElementById('relayMotor'));
ReactDOM.render(<ButtonHeater msg="Heater: Auto" />, document.getElementById('relayHeater'));
