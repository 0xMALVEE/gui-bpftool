export namespace bpfprog {
	
	export class ProgramInfo {
	    Type: string;
	    ProgramInfo?: ebpf.ProgramInfo;
	
	    static createFrom(source: any = {}) {
	        return new ProgramInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Type = source["Type"];
	        this.ProgramInfo = this.convertValues(source["ProgramInfo"], ebpf.ProgramInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

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

