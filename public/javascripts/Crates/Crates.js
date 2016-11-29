function Crates(renderer){
    this.init(
        renderer
        , "javascripts/Crates/"
        , "crates-normals.obj"
    );
}

Crates.prototype = new Collection();

Crates.prototype.process = function(engine){
};
