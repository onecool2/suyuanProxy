pragma solidity ^0.4.25;

contract CupboardContract {
/*******************************
   1 终端管理：
          云柜ID(bigint/20)；名称；设备编号(varchar/100)；CPU序列号(varchar/32)；所属合伙人(varchar/100);
          Mac地址(varchar/32)；APK版本(varchar/32)；商品表类型(smallint 1-五层   2- 13层 3-九层)；在线状态；状态(smallint)。
*******************************/
 struct yg_terminal{
      uint256  yg_id;           //  云柜ID(bigint/20)
      string   yg_name;         //  名称
      string   device_sn;       //* 设备编号(varchar/100)
      string   cpu_sn;          //* CPU序列号(varchar/32)
      string   partner;         //* 所属合伙人(varchar/100)
      bytes32  mac_addr;        //* Mac地址(varchar/32)
      string   apk_version;     //  APK版本(varchar/32)
      uint8    commodity;       //  商品表类型(smallint 1-五层   2- 13层 3-九层)
      uint8    online_status;   //  在线状态
      uint8    status;          //  状态(smallint)
  }
  mapping (uint256=> yg_terminal) private map_yg_terminal;

  function update_yg_terminal(uint256 yg_id, string yg_name, string device_sn, string cpu_sn, string partner, 
       bytes32 mac_addr, string apk_version, uint8 commodity, uint8 online_status, uint8 status) public {
    map_yg_terminal[yg_id].yg_id = yg_id;
    map_yg_terminal[yg_id].yg_name = yg_name;
    map_yg_terminal[yg_id].device_sn = device_sn;
    map_yg_terminal[yg_id].cpu_sn = cpu_sn;
    map_yg_terminal[yg_id].partner = partner;
    map_yg_terminal[yg_id].mac_addr = mac_addr;
    map_yg_terminal[yg_id].apk_version = apk_version;
    map_yg_terminal[yg_id].commodity = commodity;
    map_yg_terminal[yg_id].online_status = online_status;
    map_yg_terminal[yg_id].status = status;
    
  }

  function get_yg_terminal(uint256 yg_id) constant public returns(uint256, string, string, string, string, 
       bytes32, string, uint8, uint8, uint8 ) {
    yg_terminal memory yg = map_yg_terminal[yg_id];
    return (yg.yg_id, yg.yg_name, yg.device_sn, yg.cpu_sn, yg.partner, yg.mac_addr, yg.apk_version, yg.commodity, yg.online_status, yg.status);
  }

 
/****************************** 
  2  用户交易记录：
           用户名(varchar/50)；手机号(varchar/20)；订单编号(varchar/64)；交易号(varchar/64)；交易额(decimal/20，2)；
           创建时间(datetime)；支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)；支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)；
           计算异常(smallint计算异常0-未出错  大于0表示失败以及失败的代码)。
******************************/
 struct transaction_record{
      string   user_name;         //*  用户名(varchar/50)
      uint256  cell_number;       //*  手机号(varchar/20)
      string   order_number;      //   订单编号(varchar/64)
      string   transactoin_number;//   交易号(varchar/64)
      uint256  Transaction_amount;//   交易额(decimal/20，2)
      bytes32  create_time;       //   创建时间(datetime)
      uint8    payment_type;          //   支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)
      uint8    payment_status;           //   
      uint8    pay_exception;
  }
  
  mapping (string=> transaction_record) private map_transaction_record; 
  function update_transaction_record(string  user_name, uint256 cell_number, string order_number, string transactoin_number, uint256 Transaction_amount, 
       bytes32 create_time, uint8 payment_type, uint8 payment_status, uint8 pay_exception) public {
    map_transaction_record[user_name].user_name = user_name;
    map_transaction_record[user_name].cell_number = cell_number;
    map_transaction_record[user_name].order_number = order_number;
    map_transaction_record[user_name].transactoin_number = transactoin_number;
    map_transaction_record[user_name].Transaction_amount = Transaction_amount;
    map_transaction_record[user_name].create_time = create_time;
    map_transaction_record[user_name].payment_type = payment_type;
    map_transaction_record[user_name].payment_status = payment_status;
    map_transaction_record[user_name].pay_exception = pay_exception;
    
  }

  function get_transaction_record(string user_name) constant public returns(string, uint256, string, string, uint256, bytes32, uint8, uint8, uint8) {
    transaction_record memory transaction = map_transaction_record[user_name];
    return (transaction.user_name, transaction.cell_number, transaction.order_number, transaction.transactoin_number, 
            transaction.Transaction_amount, transaction.create_time, transaction.payment_type, transaction.payment_status, 
            transaction.pay_exception);
  }
  
/***************************** 
  3  充值记录：用户名(varchar/50)；手机号(varchar/20)；订单编号(varchar/64)；充值额(decimal/20，2)；创建时间(datetime)；支付时间(datetime)；
     支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)；支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)；微信交易ID(varchar/64)；
******************************/  
   struct recharge_record{
      string   user_name;         //* 用户名(varchar/50)
      uint256  cell_number;       //* 手机号(varchar/20)
      string   order_number;      //  订单编号(varchar/64)
      uint256  recharge_amount;   //  充值额(decimal/20，2)
      bytes32  create_time;       //  创建时间(datetime)
      bytes32  pay_time;          //  支付时间(datetime)
      uint8    pay_status;        //  支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)
      uint8    pay_type;          //  支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)
      string   wechar_id;         //* 微信交易ID(varchar/64)
  }
  
  mapping (string=> recharge_record) private map_recharge_record;
  function update_recharge_record(string  user_name, uint256 cell_number, string order_number,  uint256  recharge_amount, bytes32  create_time, 
       bytes32  pay_time, uint8 pay_status, uint8 pay_type, string wechar_id) public {
    map_recharge_record[user_name].user_name = user_name;
    map_recharge_record[user_name].cell_number = cell_number;
    map_recharge_record[user_name].order_number = order_number;
    map_recharge_record[user_name].recharge_amount = recharge_amount;
    map_recharge_record[user_name].create_time = create_time;
    map_recharge_record[user_name].pay_time = pay_time;
    map_recharge_record[user_name].pay_status = pay_status;
    map_recharge_record[user_name].pay_type = pay_type;
    map_recharge_record[user_name].wechar_id = wechar_id;
    
  }

  function get_recharge_record(string user_name) constant public returns(string, uint256, string,  uint256, bytes32, bytes32, uint8, uint8, string) {
    recharge_record memory recharge = map_recharge_record[user_name];
    return (recharge.user_name, recharge.cell_number, recharge.order_number, recharge.recharge_amount, 
            recharge.create_time, recharge.pay_time, recharge.pay_status, recharge.pay_type, 
            recharge.wechar_id);
  }
  
  /***************************** 
  4、供应链管理-大仓管理:
     仓库名称(varchar/100)；仓库地址(varchar/200)；状态；创建时间(datetime)；修改时间(datetime)
  ******************************/  
   struct big_warehouse{
      string  warehouse_name;       //   仓库名称(varchar/100)
      string  warehouse_address;    //*  仓库地址(varchar/200)
      uint8   status;               //   状态
      bytes32  create_time;          //   创建时间(datetime)
      bytes32  modify_time;          //   修改时间(datetime)
  }
  mapping (string=> big_warehouse) private map_big_warehouse;
  
  function update_big_warehouse(string  warehouse_name, string warehouse_address, uint8 status, bytes32  create_time,  bytes32  modify_time) public {
    map_big_warehouse[warehouse_name].warehouse_name = warehouse_name;
    map_big_warehouse[warehouse_name].warehouse_address = warehouse_address;
    map_big_warehouse[warehouse_name].status = status;
    map_big_warehouse[warehouse_name].create_time = create_time;
    map_big_warehouse[warehouse_name].modify_time = modify_time;
  }

  function get_big_warehouse(string warehouse_name) constant public returns (string, string, uint8, bytes32,  bytes32) {
    big_warehouse memory bigwarehouse = map_big_warehouse[warehouse_name];
    return (bigwarehouse.warehouse_name, bigwarehouse.warehouse_address, bigwarehouse.status, bigwarehouse.create_time, 
            bigwarehouse.modify_time);
  }
  
  
  /***************************** 
  5、供应链管理-云仓管理：
     仓库名称(varchar/100)；是否实仓(int仓库类型：1-真实仓；2-虚拟仓；仓库地址(varchar/200)；
     所属片区(varchar/100)；仓库管理员(varchar/100)；状态；创建时间(datetime)；修改时间(datetime)；
  ******************************/  
   struct cloud_warehouse{
      string   warehouse_name;       //  仓库名称(varchar/100)
      uint8    real;                 //  是否实仓(int仓库类型：1-真实仓；2-虚拟仓）
      string   warehouse_address;    //* 仓库地址(varchar/200)
      string   warehouse_area;       //  所属片区(varchar/100)
      string   warehouse_keeper;     //  仓库管理员(varchar/100)
      uint8    status;               //  状态
      bytes32  create_time;          //  创建时间(datetime)
      bytes32  modify_time;          //  修改时间(datetime)
  }
   mapping (string=> cloud_warehouse) private map_cloud_warehouse;
   function update_cloud_warehouse(string  warehouse_name, uint8 real, string warehouse_address, string warehouse_area, 
            string warehouse_keeper, uint8 status, bytes32  create_time, bytes32  modify_time) public {
    map_cloud_warehouse[warehouse_name].warehouse_name = warehouse_name;
    map_cloud_warehouse[warehouse_name].real = real;
    map_cloud_warehouse[warehouse_name].warehouse_address = warehouse_address;
    map_cloud_warehouse[warehouse_name].warehouse_area = warehouse_area;
    map_cloud_warehouse[warehouse_name].warehouse_keeper = warehouse_keeper;
    map_cloud_warehouse[warehouse_name].status = status;
    map_cloud_warehouse[warehouse_name].create_time = create_time;
    map_cloud_warehouse[warehouse_name].modify_time = modify_time;
  }

  function get_cloud_warehouse(string warehouse_name) constant public returns (string, uint8, string, string, string, uint8, bytes32,  bytes32) {
    cloud_warehouse memory cloudwarehouse = map_cloud_warehouse[warehouse_name];
    return (cloudwarehouse.warehouse_name, cloudwarehouse.real, cloudwarehouse.warehouse_address, cloudwarehouse.warehouse_area, 
            cloudwarehouse.warehouse_keeper, cloudwarehouse.status, cloudwarehouse.create_time, cloudwarehouse.modify_time);
  }
   
   
  /***************************** 
  6、商圈管理-商圈信息：
     商圈名称(varchar/100)；所属城市(varchar/100)；状态(int/2状态；1-有效；2-禁用；3-无效(删除)；
     创建时间(datetime)；修改时间(datetime)；
  ******************************/ 
   struct trading_area{
      string   name;            //  商圈名称(varchar/100)
      string   city;            //  所属城市(varchar/100)
      uint8    status;          //  状态(int/2状态；1-有效；2-禁用；3-无效(删除)；
      bytes32  create_time;     //  创建时间(datetime)
      bytes32  modify_time;     //  修改时间(datetime)
  }
  mapping (string=> trading_area) private map_trading_area;
  function update_trading_area(string  name, string city, uint8 status, bytes32  create_time, bytes32  modify_time) public {
    map_trading_area[name].name = name;
    map_trading_area[name].city = city;
    map_trading_area[name].status = status;
    map_trading_area[name].create_time = create_time;
    map_trading_area[name].modify_time = modify_time;
  }

  function get_trading_area(string name) constant public returns (string, string, uint8, bytes32,  bytes32) {
    trading_area memory tradingarea = map_trading_area[name];
    return (tradingarea.name, tradingarea.city, tradingarea.status, tradingarea.create_time, 
            tradingarea.modify_time);
  }
  /***************************** 
  7、商圈管理-商业片区：
     商业片区名称(varchar/100)；所属商圈(varchar/100)；
     状态(int状态；1-有效；2-禁用；3-无效(删除)；创建时间(datetime)；修改时间(datetime)；
  ******************************/ 
   struct sowntown{
      string   name;            //   商业片区名称(varchar/100)
      string   trading_area;    //   所属商圈(varchar/100)
      uint8    status;          //   状态(int状态；1-有效；2-禁用；3-无效(删除)
      bytes32  create_time;     //   创建时间(datetime)
      bytes32  modify_time;     //   修改时间(datetime)
  }
  mapping (string=> sowntown) private map_sowntown;
  function update_sowntown(string  name, string tradingarea, uint8 status, bytes32  create_time, bytes32  modify_time) public {
    map_sowntown[name].name = name;
    map_sowntown[name].trading_area = tradingarea;
    map_sowntown[name].status = status;
    map_sowntown[name].create_time = create_time;
    map_sowntown[name].modify_time = modify_time;
  }

  function get_sowntown(string name) constant public returns (string, string, uint8, bytes32,  bytes32) {
    sowntown memory sownTown = map_sowntown[name];
    return (sownTown.name, sownTown.trading_area, sownTown.status, sownTown.create_time, 
            sownTown.modify_time);
  }
  
  /***************************** 
  8、商圈管理-楼宇管理：
     楼宇名称(varchar/100)；所属片区(varchar/100)；所属云仓库(varchar/100)；
     状态(int/2状态；1-有效；2-禁用；3-无效(删除)；创建时间(datetime)；修改时间(datetime)；
  ******************************/ 
  struct building{
      string   name;                 //  楼宇名称(varchar/100)
      string   sowntown;             //  所属片区(varchar/100)
      string   cloud_warehouse_name; //  所属云仓库(varchar/100)
      uint8   status;                //  状态(int/2状态；1-有效；2-禁用；3-无效(删除)
      bytes32  create_time;          //  创建时间(datetime)
      bytes32  modify_time;          //  修改时间(datetime)
  }
  mapping (string=> building) private map_building;
  function update_building(string  name, string sownTown, uint8 status, bytes32  create_time, bytes32  modify_time) public {
    map_building[name].name = name;
    map_building[name].sowntown = sownTown;
    map_building[name].status = status;
    map_building[name].create_time = create_time;
    map_building[name].modify_time = modify_time;
  }

  function get_building(string name) constant public returns (string, string, uint8, bytes32,  bytes32) {
    building memory buildingM = map_building[name];
    return (buildingM.name, buildingM.sowntown, buildingM.status, buildingM.create_time, 
            buildingM.modify_time);
  }
  
  address public owner;
  uint public mapSize;
  
  mapping (uint=>address) public mapIndex;
  /***************************************/
  constructor () public {
      mapSize = 123;
      owner = msg.sender;
  }
}


