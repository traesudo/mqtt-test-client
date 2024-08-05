export namespace main {
	
	export class StartInput {
	    host: string;
	    username: string;
	    password: string;
	    name: string;
	    date1: string;
	    date2: string;
	    concurrencies: number;
	    topic: string;
	    message: string;
	    port: string;
	
	    static createFrom(source: any = {}) {
	        return new StartInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = source["host"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.name = source["name"];
	        this.date1 = source["date1"];
	        this.date2 = source["date2"];
	        this.concurrencies = source["concurrencies"];
	        this.topic = source["topic"];
	        this.message = source["message"];
	        this.port = source["port"];
	    }
	}

}

