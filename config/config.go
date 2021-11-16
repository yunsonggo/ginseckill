package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Listen string `json:"listen"`
	Website struct {
		WebsiteName string `json:"website_name"`
		WebsiteDesc string `json:"website_desc"`
		WebsiteMode string `json:"website_mode"`
		WebsiteLog	string `json:"website_log"`
		CorsExample string `json:"cors_example"`
		Stdout bool `json:"stdout"`
	}
	Session struct {
		SessionName string `json:"session_name"`
		SessionSecret string `json:"session_secret"`
		SessionMode string `json:"session_mode"`
		SessionSize int `json:"session_size"`
		SessionMaxage int `json:"session_maxage"`
	}
	Mysql struct {
		MysqlDriver string `json:"mysql_driver"`
		MysqlAddr string `json:"mysql_addr"`
		MysqlShow bool `json:"mysql_show"`
	}
	Redis struct {
		RedisAddr string `json:"redis_addr"`
		RedisPwd string `json:"redis_pwd"`
		RedisDb int `json:"redis_db"`
		RedisHold int `json:"redis_hold"`
	}
	Rabbitmq struct {
		RabbitmqAddr string `json:"rabbitmq_addr"`
	}
	Email struct {
		FromEmail string `json:"from_email"`
		SmtpAddr string `json:"smtp_addr"`
		SmtpPort string `json:"smtp_port"`
		SmtpPass string `json:"smtp_pass"`
		EmailSubject string `json:"email_subject"`
	}
	Jwt struct {
		JwtKey string `json:"jwt_key"`
		Issuer string `json:"issuer"`
		ExpTime int64 `json:"exp_time"`
	}
	AliSms struct {
		SmsName      string `json:"sms_name"`
		TemplateCode string `json:"template_code"`
		AppKey       string `json:"app_key"`
		AppSecret    string `json:"app_secret"`
		RegionId     string `json:"region_id"`
		TplId        string `json:"tpl_id"`
		Key          string `json:"key"`
	}
	Etcd struct {
		EtcdAddr string `json:"etcd_addr"`
		EtcdTimeout int `json:"etcd_timeout"`
	}
	Kafka struct {
		KafkaAddr string `json:"kafka_addr"`
		KafkaTopic string `json:"kafka_topic"`
	}
	Elastic struct {
		EsAddr string `json:"es_addr"`
	}
	Tail struct {
		ChannelSize int `json:"channel_size"`
	}
}

var Conf = new(Config)

func ConfigInit() {
	file,err := os.Open("config/config.json")
	if err != nil {
		err = errors.New("读取配置文件失败")
		panic(err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&Conf)
	if err != nil {
		err = errors.New("解析配置文件失败")
		panic(err.Error())
	}
	return
}