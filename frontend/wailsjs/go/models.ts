export namespace ebpf {
	
	export class ProgramInfo {
	    Type: number;
	    Tag: string;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new ProgramInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Type = source["Type"];
	        this.Tag = source["Tag"];
	        this.Name = source["Name"];
	    }
	}

}

